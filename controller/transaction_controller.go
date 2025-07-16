package controller

import (
	"net/http"
	"test-naga-exchange/model"
	"test-naga-exchange/service"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	service service.TransactionService
}

func NewTransactionHandler(s service.TransactionService) *TransactionHandler {
	return &TransactionHandler{s}
}

func (h *TransactionHandler) Get(c *gin.Context) {
	user := c.MustGet("user").(*model.User)
	result, _ := h.service.GetUserTransactions(user.ID)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Success", "data": result})
}

func (h *TransactionHandler) Process(c *gin.Context) {
	user := c.MustGet("user").(*model.User)
	var tx model.Transaction
	if err := c.ShouldBindJSON(&tx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	if err := h.service.ProcessTransaction(&tx, user.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Processed successfully"})
}
