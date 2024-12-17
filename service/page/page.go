package page

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"search-nova/internal/constant"
	"search-nova/internal/db"
	"search-nova/internal/es"
	"search-nova/internal/logger"
	"search-nova/internal/util"
	"search-nova/model/page"
)

var (
	P *pageService
)

type pageService struct {
	db *gorm.DB
}

func new() (*pageService, error) {
	return &pageService{db: db.O}, nil
}

func init() {
	var err error
	P, err = new()
	if err != nil {
		logger.L.Fatalf("Fatal error page: %v\n", err)
	}
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
