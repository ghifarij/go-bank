package migrations

import (
	"github.com/ghifarij/go-bank/helpers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserID  uint
}

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=go-bank password=112233 sslmode=disable")
	helpers.HandleErr(err)
	return db
}

func createAccounts() {
	db := connectDB()

	users := [2]User{
		{
			Username: "saiful",
			Email:    "saiful@mail.com",
		},
		{
			Username: "agus",
			Email:    "agus@mail.com",
		},
	}

	for i := 0; i < len(users); i++ {
		generatedPassowrd := helpers.HashAndSalt([]byte(users[i].Username))
		user := User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassowrd}
		db.Create(&user)

		account := Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(1000 * int(i+1)), UserID: user.ID}
		db.Create(&account)
	}
	defer db.Close()
}

func Migrate() {
	db := connectDB()
	db.AutoMigrate(&User{}, &Account{})
	defer db.Close()

	createAccounts()
}
