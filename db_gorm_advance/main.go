package main

import (
	"app/db_gorm_advance/constant"
	"app/db_gorm_advance/dbconnection"
	"app/db_gorm_advance/model"
	"app/db_gorm_advance/presenter"
	"app/db_gorm_advance/repo/mysql"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// err := createPages()

	pages, err := findAllPages()
	if err != nil {
		log.Println("findAllPages err:", err)
		return
	}

	bs, _ := json.MarshalIndent(pages, "", "  ")
	fmt.Println(string(bs))
}

func createPages() error {

	pages := []model.Page{
		{URL: "https://carlyle.vn/nguyen_minh_hai", Title: "chuyen gia lua 1 nguyen minh hai", Content: "chuyen gia lua 1 nguyen minh hai", Category: constant.EconomyCategory},
		{URL: "https://carlyle.vn/tran_thanh_binh", Title: "chuyen gia lua 2 tran thanh binh", Content: "chuyen gia lua 2 tran thanh binh", Category: constant.ScienceCategory},
	}

	pageRepo := mysql.NewPageRepository(dbconnection.GetDB())

	err := pageRepo.Create(context.Background(), pages)
	return err
}

func findAllPages() ([]presenter.Page, error) {

	pageRepo := mysql.NewPageRepository(dbconnection.GetDB())
	pages, err := pageRepo.GetAll(context.Background())
	if err != nil {
		return nil, fmt.Errorf("fail to find pages: %w", err)
	}

	result := presenter.FormPages(pages)
	return result, nil
}
