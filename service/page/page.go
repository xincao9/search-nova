package page

import (
	"errors"
	"github.com/jinzhu/gorm"
	"search-nova/internal/db"
	"search-nova/model/page"
)

var (
	P *pageService
)

func init() {
	P = new()
}

type pageService struct {
	o *gorm.DB
}

func new() *pageService {
	return &pageService{o: db.O}
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
