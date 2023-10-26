package main

import (
	"github.com/ghifarij/go-bank/api"
	"github.com/ghifarij/go-bank/database"
	_ "github.com/lib/pq"
)

func main() {
	database.InitDatabase()
	api.StartApi()
}
