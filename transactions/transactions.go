package transactions

import (
	"github.com/ghifarij/go-bank/helpers"
	"github.com/ghifarij/go-bank/interfaces"
)

// Add function create transaction
func CreateTransaction(From uint, To uint, Amount int) {
	db := helpers.ConnectDB()
	transaction := &interfaces.Transaction{From: From, To: To, Amount: Amount}
	db.Create(&transaction)

	defer db.Close()
}
