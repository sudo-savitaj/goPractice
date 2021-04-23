package dto

type TransactionRequest struct {
	toAccountNumber string
	fromAccountNumber string
	amount int
}
