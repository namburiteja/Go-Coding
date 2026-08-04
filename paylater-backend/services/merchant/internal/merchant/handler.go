package merchant

import (
	"database/sql"
	"net/http"
	"strconv"

	db "paylater/services/merchant/internal/db"
	"paylater/shared/middleware"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// RegisterRoutes mounts public merchant APIs and the internal commission endpoint
// used by Ledger over HTTP.
func RegisterRoutes(router *gin.Engine, h *Handler) {
	merchants := router.Group("/merchants")
	{
		merchants.POST("/register", h.RegisterMerchant)
		merchants.POST("/login", h.LoginMerchant)

		merchants.GET(
			"",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			h.GetAllMerchants,
		)

		merchants.GET(
			"/me",
			middleware.AuthMiddleware(),
			middleware.MerchantOnly(),
			h.GetMyProfile,
		)

		merchants.PUT(
			"/me",
			middleware.AuthMiddleware(),
			middleware.MerchantOnly(),
			h.UpdateMyProfile,
		)

		merchants.PUT(
			"/:id/commission",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			h.UpdateMerchantCommission,
		)

		merchants.GET(
			"/:id",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			h.GetMerchantByID,
		)

		merchants.PUT(
			"/:id",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			h.UpdateMerchant,
		)

		merchants.DELETE(
			"/:id",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			h.DeleteMerchant,
		)
	}

	// Service-to-service: Ledger reads commission without end-user JWT.
	router.GET("/internal/merchants/:id/commission", h.GetCommissionInternal)
}

func (h *Handler) RegisterMerchant(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.RegisterMerchant(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Merchant registered successfully"})
}

func (h *Handler) LoginMerchant(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.LoginMerchant(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: token})
}

func (h *Handler) GetMyProfile(c *gin.Context) {
	userID := c.MustGet("userID").(int32)

	merchant, err := h.service.GetMyProfile(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, merchant)
}

func (h *Handler) UpdateMyProfile(c *gin.Context) {
	userID := c.MustGet("userID").(int32)

	var req db.UpdateMerchantParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.UpdateMyProfile(c.Request.Context(), userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

func (h *Handler) GetMerchantByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Merchant ID"})
		return
	}

	merchant, err := h.service.GetMerchantByID(c.Request.Context(), int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Merchant Not Found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, merchant)
}

func (h *Handler) GetAllMerchants(c *gin.Context) {
	merchants, err := h.service.GetAllMerchants(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, merchants)
}

func (h *Handler) UpdateMerchantCommission(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Merchant ID"})
		return
	}

	var req db.UpdateMerchantCommissionParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ID = int32(id)

	err = h.service.UpdateMerchantCommission(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Merchant Commission Updated Successfully"})
}

func (h *Handler) UpdateMerchant(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Merchant ID"})
		return
	}

	var req db.UpdateMerchantParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ID = int32(id)

	err = h.service.UpdateMerchant(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Merchant Updated Successfully"})
}

func (h *Handler) DeleteMerchant(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Merchant ID"})
		return
	}

	err = h.service.DeleteMerchant(c.Request.Context(), int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Merchant Deleted Successfully"})
}

// GetCommissionInternal is called by Ledger over REST (not exposed via gateway).
func (h *Handler) GetCommissionInternal(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Merchant ID"})
		return
	}

	merchant, err := h.service.GetMerchantByID(c.Request.Context(), int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "merchant not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var percentage *string
	if merchant.CommissionPercentage.Valid {
		percentage = &merchant.CommissionPercentage.String
	}

	c.JSON(http.StatusOK, CommissionResponse{
		ID:                   merchant.ID,
		CommissionPercentage: percentage,
	})
}
