package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Page struct {
	URL     string
	Title   string
	Content string
}

func (p *Page) ToString() []string {
	result := []string{p.URL, p.Title, p.Content}
	return result
}

func main() {
	num := 0
	pages := []Page{}
	c := colly.NewCollector(
		colly.AllowedDomains("umcclinic.com.vn"),
		colly.CacheDir("/Users/mfv-computer-0140/dhyd_cache"),
	)

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnHTML("html", func(e *colly.HTMLElement) {
		num++
		content, err := e.DOM.Html()
		if err != nil {
			log.Println(err)
		}
		// content := string(e.Response.Body)
		title := e.ChildText("title")
		page := Page{
			URL:     e.Request.URL.String(),
			Title:   title,
			Content: content,
		}

		pages = append(pages, page)

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://umcclinic.com.vn/")

	fmt.Println("total web page:", num, len(pages))
	// writeToCSV(pages)
	writeToDB(pages)
}

func writeToCSV(pages []Page) {
	f, err := os.Create("pages.csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	for i := range pages {
		if err := w.Write(pages[i].ToString()); err != nil {
			log.Fatalln("error write csv", err)
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}

func writeToDB(pages []Page) {
	dbname := "scraper"
	user := "admin"
	password := "moneyforward123"
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:30001)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatalln("failed to connect db", err)
	}

	result := db.Create(pages)

	if result.Error != nil {
		fmt.Println(result.Error)
	} else {
		fmt.Println("affected rows", result.RowsAffected)
	}
}
