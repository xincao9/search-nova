package page

import (
	"search-nova/internal/db"
)

type Page struct {
	db.Model
	Md5      string `json:"md5" gorm:"md5"`
	Url      string `json:"url" gorm:"column:url"`
	Title    string `json:"title" gorm:"column:title" goquery:"title"`
	Describe string `json:"describe" gorm:"column:describe" goquery:"h1"`
	Keywords string `json:"keywords" gorm:"column:keywords"`
	Content  string `json:"content" gorm:"-"`
	Status   int    `json:"status" gorm:"column:status"`
}

type SearchRequest struct {
	Query struct {
		Match struct {
			Content string `json:"content"`
		} `json:"match"`
	} `json:"query"`
}

type SearchResponse struct {
	Hits struct {
		Hits []struct {
			Source struct {
				Id int64 `json:"id"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
