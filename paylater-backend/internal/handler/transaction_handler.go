package handler

import (
	"database/sql"
	"net/http"
	"strconv"
	// "fmt"

	// db "paylater-backend/internal/db"
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


func (h *TransactionHandler) Purchase(c *gin.Context) {

	var req dto.PurchaseRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.MustGet("userID").(int32)

	err := h.service.Purchase(
		c.Request.Context(),
		userID,
		req,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Purchase successful",
	})
}

func (h *TransactionHandler) Payback(c *gin.Context) {

	var req dto.PaybackRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.MustGet("userID").(int32)

	err := h.service.Payback(
		c.Request.Context(),
		userID,
		req,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Payment successful",
	})
}

func (h *TransactionHandler) GetMyTransactions(c *gin.Context) {

	userID := c.MustGet("userID").(int32)

	transactions, err := h.service.GetTransactionsByCustomerID(
		c.Request.Context(),
		userID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *TransactionHandler) GetMerchantTransactions(c *gin.Context) {

	merchantID := c.MustGet("userID").(int32)

	transactions, err := h.service.GetTransactionsByMerchantID(
		c.Request.Context(),
		sql.NullInt32{
			Int32: merchantID,
			Valid: true,
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, transactions)
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

