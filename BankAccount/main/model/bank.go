package model


type BankInfo struct {
	name string
	branch string
	ifscCode string
	address string
}

type Bank struct {
	*BankInfo
	employees []Employee
	accounts []Account
}