package model

type AccountType int
const (
	Savings AccountType = 0 + iota
	Current
	Overdraft
)

type TransactionType int
const (
	Debit TransactionType = 0 + iota
	Credit
)

type TransactionInfo struct {
	id string //random
	reason    string
	amount    int
	transType TransactionType
}

type Account struct {
	transactions  []TransactionInfo
	customerInfo  CustomerInfo
	accountType   AccountType
	bankInfo      BankInfo
	accountNumber int
	amount        int
}

func NewAccount(accountType AccountType, accountNumber int, amount int) *Account {
	return &Account{
		accountType:   accountType,
		accountNumber: accountNumber,
		amount:        amount,
	}
}

func (this *Account) Deposit(amount int,reason string) {
	this.amount += amount
	this.transactions = append(this.transactions,TransactionInfo{
		reason:    reason,
		amount:    amount,
		transType: Credit,
	})
}

func (this *Account) Withdraw(amount int,reason string) {
	this.amount -= amount
	this.transactions = append(this.transactions,TransactionInfo{
		reason:    reason,
		amount:    amount,
		transType: Debit,
	})
}
