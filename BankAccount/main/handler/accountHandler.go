package handler

import (
	"Hackathon/BankAccount/main/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type accountHandler struct {
	accountService service.AccountService
}

func NewAccountHandler(accountService service.AccountService) *accountHandler{
	return &accountHandler{
		accountService: accountService,
	}
}

func (handler *accountHandler) Get(context *gin.Context) {
	context.JSON(http.StatusOK, "Account info...here it goes")
}

func (handler *accountHandler) Create(context *gin.Context) {
	context.JSON(http.StatusOK, "Account info...here it goes")
}
