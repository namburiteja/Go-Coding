package handler

import (
	"database/sql"
	"net/http"
	"strconv"
	// "fmt"

	db "paylater-backend/internal/db"
	"paylater-backend/internal/service"
	"paylater-backend/internal/dto"
	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(service *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		service: service,
	}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {

	var req dto.CreateTransactionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	params := db.CreateTransactionParams{
		CustomerID:      req.CustomerID,
		TransactionType: db.TransactionsTransactionType(req.TransactionType),
		Amount:          req.Amount,
	}

	if req.MerchantID != nil {

		params.MerchantID = sql.NullInt32{
			Int32: *req.MerchantID,
			Valid: true,
		}

	}

	err := h.service.CreateTransaction(c.Request.Context(), params)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Transaction Created Successfully",
	})

}

func (h *TransactionHandler) GetAllTransactions(c *gin.Context) {

	transactions, err := h.service.GetAllTransactions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *TransactionHandler) GetTransactionsByCustomerID(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Customer ID",
		})
		return
	}

	transactions, err := h.service.GetTransactionsByCustomerID(c, int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *TransactionHandler) GetTransactionsByMerchantID(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Merchant ID",
		})
		return
	}

	transactions, err := h.service.GetTransactionsByMerchantID(c, sql.NullInt32{
		Int32: int32(id),
		Valid: true,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

