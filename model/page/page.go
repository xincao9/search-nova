package page

import (
	"search-nova/internal/db"
)

type Page struct {
	db.Model
	Url      string `json:"url" gorm:"column:url"`
	Title    string `json:"title" gorm:"column:title" goquery:"title"`
	Describe string `json:"describe" gorm:"column:describe" goquery:"h1"`
	Keywords string `json:"keywords" gorm:"column:keywords"`
	Content  string `json:"content"`
}
