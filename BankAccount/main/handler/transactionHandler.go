package handler

import (
	"Hackathon/BankAccount/main/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type transactiontHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) *transactiontHandler{
	return &transactiontHandler{
		transactionService: service.TransactionService{},
	}
}

func (handler *transactiontHandler) Get(context *gin.Context) {
	context.JSON(http.StatusOK, "Account info...here it goes")
}

func (handler *transactiontHandler) Create(context *gin.Context) {
	context.JSON(http.StatusOK, "Account info...here it goes")
}
