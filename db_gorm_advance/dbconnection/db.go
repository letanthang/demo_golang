package dbconnection

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB
	once     sync.Once
)

func GetDB() *gorm.DB {
	once.Do(func() {
		dbname := "scraper"
		user := "admin"
		password := "moneyforward123"
		dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:30001)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, dbname)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		instance = db
	})

	return instance

}
