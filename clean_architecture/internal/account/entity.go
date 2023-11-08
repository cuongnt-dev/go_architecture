package account

import "time"

type TransactionType string

const (
	Send TransactionType = "send"
)

type Transaction struct {
	From            string `json:"userId"`
	TransactionType string `json:"transactionType"`
	Amount          int    `json:"amount"`
	To              string `json:"to"`
	CreatedAt       time.Time
}

type Account struct {
	UserName    string        `json:"userName"`
	Address     string        `json:"address"`
	Transaction []Transaction `json:"transaction"`
}
