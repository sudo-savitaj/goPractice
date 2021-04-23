package handler

import (
"github.com/gin-gonic/gin"
"net/http"
)

type healthHandler struct {
}

func NewHealthHandler() *healthHandler {
	return &healthHandler{}
}

func (handler *healthHandler) Get(context *gin.Context) {
	context.JSON(http.StatusOK, "I am running...")
}

