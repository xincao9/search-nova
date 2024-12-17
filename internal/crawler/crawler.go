package crawler

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/panjf2000/ants/v2"
	"net/url"
	"regexp"
	"search-nova/internal/chromium"
	"search-nova/internal/constant"
	"search-nova/internal/es"
	"search-nova/internal/logger"
	"search-nova/internal/util"
	mp "search-nova/model/page"
	"search-nova/service/page"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	C *crawler
)

type crawler struct {
	running atomic.Bool
	pool    *ants.Pool
}

func init() {
	var err error
	C, err = new()
	if err != nil {
		logger.L.Fatalf("Fatal error crawler: %v\n", err)
	}
}
func new() (*crawler, error) {
	c := &crawler{}
	var err error
	c.running.Store(false)
	c.pool, err = ants.NewPool(4)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *crawler) Refresh() error {
	maxId, err := page.P.MaxId()
	if err != nil {
		return err
	}
	var id int64 = 1
	var wg sync.WaitGroup
	for ; id <= maxId; id++ {
		logger.L.Infof("crawler.Refresh(): %d/%d\n", id, maxId)
		p, err := page.P.GetPageById(id)
		if err != nil {
			logger.L.Errorf("page.GetPageById(%d) err %v\n", id, err)
			continue
		}
		if p == nil {
			continue
		}
		if p.Status != constant.NewStatus {
			continue
		}
		wg.Add(1)
		err = c.pool.Submit(func() {
			defer wg.Done()
			err = c.TextAnalysis(p)
			if err == nil {
				return
			}
			logger.L.Errorf("crawler.TextAnalysis(%s) err %v\n", p.Url, err)
			p.Status = constant.FailureStatus
			err = page.P.Save(p)
			if err != nil {
				logger.L.Errorf("page.Save(%v) err %v\n", p, err)
			}
		})
	}
	wg.Wait()
	return nil
}

func (c *crawler) Run() {
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		if !c.running.CompareAndSwap(false, true) {
			logger.L.Infoln("crawler.Refresh() running")
			continue
		}
		go func() {
			defer c.running.Store(false)
			err := c.Refresh()
			if err != nil {
				logger.L.Errorf("crawler.Refresh() err: %v\n", err)
			}
		}()
	}
}

func (c *crawler) TextAnalysis(p *mp.Page) error {
	urlS := p.Url
	urlS = strings.TrimSpace(urlS)
	urlO, err := url.Parse(urlS)
	if err != nil {
		return err
	}
	reader, err := chromium.C.Html(urlS)
	if err != nil {
		return err
	}
	if reader == nil {
		return nil
	}
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return err
	}
	doc.Find("title").Each(func(index int, item *goquery.Selection) {
		p.Title = item.Text()
	})
	doc.Find("meta[name='description']").Each(func(index int, item *goquery.Selection) {
		if _, exists := item.Attr("content"); exists {
			p.Describe = item.AttrOr("content", "")
		}
	})
	doc.Find("meta[name='keywords']").Each(func(index int, item *goquery.Selection) {
		if _, exists := item.Attr("content"); exists {
			p.Keywords = item.AttrOr("content", "")
		}
	})
	p.Content = c.extractText(doc)
	if p.Keywords == "" && p.Content != "" {
		// TODO 提取关键词
	}
	p.Status = constant.SuccessStatus
	err = page.P.Save(p)
	if err != nil {
		return err
	}
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}
	err = es.E.IndexDoc(data)
	if err != nil {
		return err
	}
	doc.Find("a").Each(func(index int, s *goquery.Selection) {
		val, exists := s.Attr("href")
		if !exists {
			return
		}
		url1, err := url.Parse(val)
		if err != nil {
			logger.L.Errorf("url.Parse err: %v\n", err)
			return
		}
		np := &mp.Page{}
		if url1.Scheme == "" {
			np.Url = urlO.ResolveReference(url1).String()
		} else if url1.Scheme == "http" || url1.Scheme == "https" {
			np.Url = url1.String()
		} else {
			return
		}
		np.Status = constant.NewStatus
		np.Md5 = util.Md5([]byte(np.Url))
		err = page.P.Save(np)
		if err != nil {
			logger.L.Errorf("page.Save err: %v\n", err)
			return
		}
	})
	return nil
}

func (c *crawler) extractText(doc *goquery.Document) string {
	var builder strings.Builder
	chinesePattern := regexp.MustCompile(`\p{Han}+`)
	englishPattern := regexp.MustCompile(`\b[a-zA-Z]+\b`)
	doc.Find("*").Each(func(index int, s *goquery.Selection) {
		text := s.Text()
		chineseMatches := chinesePattern.FindAllString(text, -1)
		for _, match := range chineseMatches {
			builder.WriteString(match)
		}
		englishMatches := englishPattern.FindAllString(text, -1)
		for _, match := range englishMatches {
			builder.WriteString(match)
			builder.WriteRune(' ')
		}
	})
	return builder.String()
}