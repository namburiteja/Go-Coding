package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	db "paylater-backend/internal/db"
	"paylater-backend/internal/service"
	"paylater-backend/internal/dto"

	"github.com/gin-gonic/gin"
)

type MerchantHandler struct {
	service *service.MerchantService
}

func NewMerchantHandler(service *service.MerchantService) *MerchantHandler {
	return &MerchantHandler{
		service: service,
	}
}

// Create Merchant
func (h *MerchantHandler) RegisterMerchant(c *gin.Context) {

	var req dto.MerchantRegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.service.RegisterMerchant(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Merchant registered successfully",
	})
}

//login merchant
func (h *MerchantHandler) LoginMerchant(c *gin.Context) {

	var req dto.MerchantLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := h.service.LoginMerchant(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.MerchantLoginResponse{
		Token: token,
	})
}

func (h *MerchantHandler) GetMyProfile(c *gin.Context) {

	userID := c.MustGet("userID").(int32)

	merchant, err := h.service.GetMyProfile(
		c.Request.Context(),
		userID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, merchant)
}

func (h *MerchantHandler) UpdateMyProfile(c *gin.Context) {

	userID := c.MustGet("userID").(int32)

	var req db.UpdateMerchantParams

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.service.UpdateMyProfile(
		c.Request.Context(),
		userID,
		req,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
	})
}



// Get Merchant By ID
func (h *MerchantHandler) GetMerchantByID(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Merchant ID",
		})
		return
	}

	merchant, err := h.service.GetMerchantByID(c, int32(id))
	if err != nil {

		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Merchant Not Found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, merchant)
}

// Get All Merchants
func (h *MerchantHandler) GetAllMerchants(c *gin.Context) {

	merchants, err := h.service.GetAllMerchants(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, merchants)
}

// Update Merchant Commission
func (h *MerchantHandler) UpdateMerchantCommission(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Merchant ID",
		})
		return
	}

	var req db.UpdateMerchantCommissionParams

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	req.ID = int32(id)

	err = h.service.UpdateMerchantCommission(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Merchant Commission Updated Successfully",
	})
}

func (h *MerchantHandler) UpdateMerchant(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Merchant ID",
		})
		return
	}

	var req db.UpdateMerchantParams

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	req.ID = int32(id)

	err = h.service.UpdateMerchant(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Merchant Updated Successfully",
	})
}

func (h *MerchantHandler) DeleteMerchant(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Merchant ID",
		})
		return
	}

	err = h.service.DeleteMerchant(c, int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Merchant Deleted Successfully",
	})
}