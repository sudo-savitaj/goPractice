package model

type Sender interface {
	Withdraw(ammount int, reason string)
}

type Receiver interface{
	Deposit(ammount int, reason string)
}

type TransactionStatus int
const(
	Initiated TransactionStatus = 0 + iota
	Successfull
	Failure
)

type Transaction struct {
	sender Sender
	receiver Receiver
	ammount int
	reason string
	status TransactionStatus
}

func NewTransaction(sender Sender, receiver Receiver, ammount int,reason string ,status TransactionStatus) *Transaction{
	return &Transaction{
		sender:   sender,
		receiver: receiver,
		ammount:  ammount,
		reason:reason,
		status:status,
	}
}

func (this *Transaction) Process(){
	if this.status == Successfull {
		return
	}
	this.sender.Withdraw(this.ammount,this.reason)
	this.receiver.Deposit(this.ammount,this.reason)
	this.status = Successfull
}