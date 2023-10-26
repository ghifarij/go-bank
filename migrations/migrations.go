package migrations

import (
	"github.com/ghifarij/go-bank/helpers"
	"github.com/ghifarij/go-bank/interfaces"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createAccounts() {
	db := helpers.ConnectDB()

	users := [2]interfaces.User{
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
		user := interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassowrd}
		db.Create(&user)

		account := interfaces.Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(1000 * int(i+1)), UserID: user.ID}
		db.Create(&account)
	}
	defer db.Close()
}

func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	db := helpers.ConnectDB()
	db.AutoMigrate(&User, &Account)
	defer db.Close()

	createAccounts()
}
