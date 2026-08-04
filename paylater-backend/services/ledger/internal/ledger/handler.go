package ledger

import (
	"database/sql"
	"net/http"
	"strconv"

	"paylater/shared/middleware"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// RegisterRoutes mounts ledger-owned public APIs (same paths as before extraction).
func RegisterRoutes(router *gin.Engine, h *Handler) {
	router.POST(
		"/customers/purchase",
		middleware.AuthMiddleware(),
		middleware.CustomerOnly(),
		h.Purchase,
	)
	router.POST(
		"/customers/payback",
		middleware.AuthMiddleware(),
		middleware.CustomerOnly(),
		h.Payback,
	)
	router.GET(
		"/customers/me/transactions",
		middleware.AuthMiddleware(),
		middleware.CustomerOnly(),
		h.GetMyTransactions,
	)
	router.GET(
		"/merchants/me/transactions",
		middleware.AuthMiddleware(),
		middleware.MerchantOnly(),
		h.GetMerchantTransactions,
	)
	router.GET(
		"/transactions",
		middleware.AuthMiddleware(),
		middleware.AdminOnly(),
		h.GetAllTransactions,
	)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	c.JSON(http.StatusOK, transactions)
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

	c.JSON(http.StatusOK, transactions)
}

func (h *Handler) GetAllTransactions(c *gin.Context) {
	transactions, err := h.service.GetAllTransactions(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *Handler) GetTransactionsByCustomerID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Customer ID"})
		return
	}

	transactions, err := h.service.GetTransactionsByCustomerID(c.Request.Context(), int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *Handler) GetTransactionsByMerchantID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Merchant ID"})
		return
	}

	transactions, err := h.service.GetTransactionsByMerchantID(c.Request.Context(), sql.NullInt32{
		Int32: int32(id),
		Valid: true,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
