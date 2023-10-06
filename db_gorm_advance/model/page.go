package model

import "app/db_gorm_advance/constant"

type Page struct {
	URL      string
	Title    string
	Content  string
	Category constant.PageCagegory
}

func (p *Page) TableName() string {
	return "pages"
}
