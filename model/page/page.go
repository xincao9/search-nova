package page

import (
	"search-nova/internal/db"
)

type Page struct {
	db.Model
	Url      string `json:"url" gorm:"column:url"`
	Title    string `json:"title" gorm:"column:title"`
	Describe string `json:"describe" gorm:"column:describe"`
	Keywords string `json:"keywords" gorm:"column:keywords"`
}
