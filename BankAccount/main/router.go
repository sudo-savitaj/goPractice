package main

import (
	"Hackathon/BankAccount/main/config"
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
	"Hackathon/BankAccount/main/handler"
	"Hackathon/BankAccount/main/service"
)

const (
	HealthEndPoint = "health"
	Account = "account"
	Transaction = "transaction"
)

func RegisterRoutes(config *config.Config, engine *gin.Engine) *gin.Engine {
	accountService := service.AccountService{}
	transactionService := service.TransactionService{}

	engine.GET(HealthEndPoint, handler.NewHealthHandler().Get).
		GET(Account, handler.NewAccountHandler(accountService).Get).
		POST(Account, handler.NewAccountHandler(accountService).Create).
		GET(Transaction, handler.NewTransactionHandler(transactionService).Get).
		POST(Transaction, handler.NewTransactionHandler(transactionService).Create)

	return engine
}

func main() {
	startApp()
}

func startApp() {
	api := gin.New()
	config := config.LoadConfig()
	time.Local = time.UTC

	address := config.Host + ":" + config.Port
	fmt.Println("API available at " + address)

	runServer(address, config, api)
}

func runServer(address string, config *config.Config, apiEngine *gin.Engine) {
	RegisterRoutes(config, apiEngine).
		Run(address)
}
