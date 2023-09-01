package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Page struct {
	URL     string
	Title   string
	Content string
}

func (p *Page) TableName() string {
	return "pages"
}

func main() {
	dbname := "scraper"
	user := "admin"
	password := "moneyforward123"
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:30001)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("failed to connect db", err)
	}

	pages := []Page{
		{URL: "https://carlyle.vn/nguyen_minh_hai", Title: "chuyen gia lua 1 nguyen minh hai", Content: "chuyen gia lua 1 nguyen minh hai"},
		{URL: "https://carlyle.vn/tran_thanh_binh", Title: "chuyen gia lua 2 tran thanh binh", Content: "chuyen gia lua 2 tran thanh binh"},
	}

	result := db.Create(pages)

	if result.Error != nil {
		fmt.Println(result.Error)
	} else {
		fmt.Println("affected rows", result.RowsAffected)
	}

}
