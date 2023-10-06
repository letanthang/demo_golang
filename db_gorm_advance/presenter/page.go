package presenter

import "app/db_gorm_advance/model"

type Page struct {
	URL      string
	Title    string
	Content  string
	Category string
}

func FormPage(page model.Page) Page {
	obj := Page{
		URL:      page.URL,
		Title:    page.Title,
		Content:  page.Content,
		Category: page.Category.String(),
	}
	return obj
}

func FormPages(pages []model.Page) []Page {
	objs := make([]Page, len(pages))

	for i := range pages {
		objs[i] = FormPage(pages[i])
	}
	return objs
}
