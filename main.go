package main

import (
	"github.com/ghifarij/go-bank/api"
	_ "github.com/lib/pq"
)

func main() {
	// migrations.Migrate()
	api.StartApi()
}
