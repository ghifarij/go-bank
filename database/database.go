package database

import (
	"github.com/ghifarij/go-bank/helpers"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	database, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=go-bank password=112233 sslmode=disable")
	helpers.HandleErr(err)

	database.DB().SetMaxIdleConns(20)
	database.DB().SetMaxOpenConns(200)
	DB = database
}
