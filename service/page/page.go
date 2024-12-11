package page

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	"net/http"
	"search-nova/internal/db"
	"search-nova/internal/logger"
	"search-nova/model/page"
	"time"
)

var (
	P *pageService
)

func init() {
	P = new()
	go func() {
		ticker := time.NewTicker(time.Hour)
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
		// TODO 已经下载过的，不再下载
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

func (ps *pageService) TextAnalysis(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	p := &page.Page{
		Url: url,
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
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
	// TODO 向下遍历
	return ps.Save(p)
}

func (ps *pageService) Save(p *page.Page) error {
	op, err := ps.GetPageByUrl(p.Url)
	if err != nil {
		return err
	}
	p.Id = op.Id
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
