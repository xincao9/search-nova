package page

import (
	"bytes"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/html/charset"
	"net/http"
	"net/url"
	"search-nova/internal/db"
	"search-nova/internal/logger"
	"search-nova/model/page"
	"strings"
	"time"
)

var (
	P *pageService
)

func init() {
	P = new()
	go func() {
		ticker := time.NewTicker(time.Minute)
		for range ticker.C {
			err := P.Refresh()
			if err != nil {
				logger.L.Errorf("page.Refresh err: %v\n", err)
			}
		}
	}()
}

func (ps *pageService) Refresh() error {
	maxId, err := ps.MaxId()
	if err != nil {
		return err
	}
	var id int64 = 1
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
		if p.Title != "" {
			continue
		}
		err = P.TextAnalysis(p.Url)
		if err != nil {
			logger.L.Errorf("page.TextAnalysis(%s) err %v\n", p.Url, err)
			continue
		}
	}
	return nil
}

type pageService struct {
	o *gorm.DB
}

func new() *pageService {
	return &pageService{o: db.O}
}

func (ps *pageService) TextAnalysis(urlS string) error {
	urlO, err := url.Parse(urlS)
	if err != nil {
		return err
	}
	resp, err := http.Get(urlS)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	ct := resp.Header.Get("Content-Type")
	if ct != "" && !strings.Contains(ct, "text/html") {
		logger.L.Infof("url: %s content-type is %s\n", urlS, ct)
		return nil
	}
	reader, err := charset.NewReader(resp.Body, ct)
	if err != nil {
		return err
	}
	p := &page.Page{
		Url: urlS,
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
	var buf bytes.Buffer
	doc.Find("*").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		for _, r := range text {
			if buf.Len() >= 10240 {
				return
			}
			buf.WriteRune(r)
		}
	})
	p.Content = buf.String()
	if p.Keywords == "" && p.Content != "" {
		// TODO 提取关键词
	}
	err = ps.Save(p)
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
		} else {
			np.Url = url1.String()
		}
		err = ps.Save(np)
		if err != nil {
			logger.L.Errorf("page.Save err: %v\n", err)
			return
		}
	})
	return nil
}

func (ps *pageService) Save(p *page.Page) error {
	op, err := ps.GetPageByUrl(p.Url)
	if err != nil {
		return err
	}
	if op != nil {
		p.Id = op.Id
	}
	err = ps.o.Save(p).Error
	return err
}

func (ps *pageService) GetPageByUrl(url string) (*page.Page, error) {
	p := &page.Page{}
	err := ps.o.Where("`url`=?", url).First(p).Error
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
	err := ps.o.Where("`id`=?", id).First(p).Error
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
	err := ps.o.Order("`id` desc").First(p).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return p.Id, nil
}