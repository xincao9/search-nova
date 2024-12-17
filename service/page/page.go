package page

import (
	"encoding/json"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	"github.com/panjf2000/ants/v2"
	"net/url"
	"regexp"
	"search-nova/internal/constant"
	"search-nova/internal/db"
	"search-nova/internal/logger"
	"search-nova/internal/util"
	"search-nova/model/page"
	"search-nova/service/chromium"
	"search-nova/service/es"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	P *pageService
)

type pageService struct {
	db      *gorm.DB
	running atomic.Bool
	pool    *ants.Pool
}

func new() (*pageService, error) {
	ps := &pageService{db: db.O}
	var err error
	ps.running.Store(false)
	ps.pool, err = ants.NewPool(4)
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func init() {
	var err error
	P, err = new()
	if err != nil {
		logger.L.Fatalf("Fatal error page: %v\n", err)
	}
	go func() {
		ticker := time.NewTicker(time.Minute)
		for range ticker.C {
			if !P.running.CompareAndSwap(false, true) {
				logger.L.Infoln("page.Refresh running")
				continue
			}
			go func() {
				defer P.running.Store(false)
				err := P.Refresh()
				if err != nil {
					logger.L.Errorf("page.Refresh err: %v\n", err)
				}
			}()
		}
	}()
}

func (ps *pageService) Refresh() error {
	maxId, err := ps.MaxId()
	if err != nil {
		return err
	}
	var id int64 = 1
	var wg sync.WaitGroup
	for ; id <= maxId; id++ {
		logger.L.Infof("page.Refresh: %d/%d\n", id, maxId)
		p, err := P.GetPageById(id)
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
		err = ps.pool.Submit(func() {
			defer wg.Done()
			err = P.TextAnalysis(p)
			if err == nil {
				return
			}
			logger.L.Errorf("page.TextAnalysis(%s) err %v\n", p.Url, err)
			p.Status = constant.FailureStatus
			err = ps.Save(p)
			if err != nil {
				logger.L.Errorf("page.Save(%v) err %v\n", p, err)
			}
		})
	}
	wg.Wait()
	return nil
}

func (ps *pageService) TextAnalysis(p *page.Page) error {
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
	p.Content = ps.extractText(doc)
	if p.Keywords == "" && p.Content != "" {
		// TODO 提取关键词
	}
	p.Status = constant.SuccessStatus
	err = ps.Save(p)
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
		np := &page.Page{}
		if url1.Scheme == "" {
			np.Url = urlO.ResolveReference(url1).String()
		} else if url1.Scheme == "http" || url1.Scheme == "https" {
			np.Url = url1.String()
		} else {
			return
		}
		np.Status = constant.NewStatus
		np.Md5 = util.Md5([]byte(np.Url))
		err = ps.Save(np)
		if err != nil {
			logger.L.Errorf("page.Save err: %v\n", err)
			return
		}
	})
	return nil
}

func (ps *pageService) extractText(doc *goquery.Document) string {
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

func (ps *pageService) Save(p *page.Page) error {
	op, err := ps.GetPageByUrl(p.Url)
	if err != nil {
		return err
	}
	if op != nil {
		if op.Status != constant.NewStatus {
			// 老数据非新建的状态，不能再更改
			return nil
		}
		p.Id = op.Id
	}
	err = ps.db.Save(p).Error
	return err
}

func (ps *pageService) Match(text string) ([]*page.Page, error) {
	var sr page.SearchRequest
	sr.Query.Match.Content = text
	body, err := json.Marshal(sr)
	if err != nil {
		return nil, err
	}
	searchResponse, err := es.E.Search(body)
	if err != nil {
		return nil, err
	}
	var pages []*page.Page
	for _, hit := range searchResponse.Hits.Hits {
		if hit.Source.Id <= 0 {
			continue
		}
		p, err := ps.GetPageById(hit.Source.Id)
		if err != nil {
			continue
		}
		if p == nil {
			continue
		}
		pages = append(pages, p)
	}
	return pages, nil
}

func (ps *pageService) GetPageByUrl(url string) (*page.Page, error) {
	md5 := util.Md5([]byte(url))
	p := &page.Page{}
	err := ps.db.Where("`md5`=?", md5).First(p).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (ps *pageService) GetPageById(id int64) (*page.Page, error) {
	p := &page.Page{}
	err := ps.db.Where("`id`=?", id).First(p).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (ps *pageService) MaxId() (int64, error) {
	p := &page.Page{}
	err := ps.db.Order("`id` desc").First(p).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return p.Id, nil
}
