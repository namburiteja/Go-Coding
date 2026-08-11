package ledger

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Purchase(c *gin.Context) {
	var req PurchaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("userID").(int32)

	err := h.service.Purchase(c.Request.Context(), userID, req)
	if err != nil {
		c.JSON(purchaseErrorStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Purchase successful"})
}

func (h *Handler) Payback(c *gin.Context) {
	var req PaybackRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("userID").(int32)

	err := h.service.Payback(c.Request.Context(), userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment successful"})
}

func (h *Handler) GetMyTransactions(c *gin.Context) {
	userID := c.MustGet("userID").(int32)

	transactions, err := h.service.GetTransactionsByCustomerID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, toTransactionResponses(transactions))
}

func (h *Handler) GetMerchantTransactions(c *gin.Context) {
	merchantID := c.MustGet("userID").(int32)

	transactions, err := h.service.GetTransactionsByMerchantID(
		c.Request.Context(),
		sql.NullInt32{Int32: merchantID, Valid: true},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, toTransactionResponses(transactions))
}

func (h *Handler) GetAllTransactions(c *gin.Context) {
	transactions, err := h.service.GetAllTransactions(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, toTransactionResponses(transactions))
}

func (h *Handler) GetMerchantFeesInternal(c *gin.Context) {
	fees, err := h.service.GetMerchantFeeTotals(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, fees)
}

func purchaseErrorStatus(err error) int {
	msg := err.Error()
	if msg == "customer is blocked" ||
		strings.Contains(msg, "payment due date crossed") {
		return http.StatusForbidden
	}
	return http.StatusBadRequest
}
