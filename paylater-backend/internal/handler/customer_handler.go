package handler

import (
	"net/http"
	"strconv"
	"database/sql"
	db "paylater-backend/internal/db"
	"paylater-backend/internal/service"
	"paylater-backend/internal/dto"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	service *service.CustomerService
}

func NewCustomerHandler(service *service.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		service: service,
	}
}

func (h *CustomerHandler) RegisterCustomer(c *gin.Context) {

	var req dto.CustomerRegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.service.RegisterCustomer(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Customer registered successfully",
	})
}

func (h *CustomerHandler) LoginCustomer(c *gin.Context) {

	var req dto.CustomerLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := h.service.LoginCustomer(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.CustomerLoginResponse{
		Token: token,
	})
}

func (h *CustomerHandler) GetMyProfile(c *gin.Context) {

	userID := c.MustGet("userID").(int32)

	customer, err := h.service.GetMyProfile(
		c.Request.Context(),
		userID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (h *CustomerHandler) UpdateMyProfile(c *gin.Context) {

	userID := c.MustGet("userID").(int32)

	var req db.UpdateCustomerParams

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

// GET /customers/:id
func (h *CustomerHandler) GetCustomerByID(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Customer ID"})
		return
	}

	customer, err := h.service.GetCustomerByID(c, int32(id))
	if err != nil {

		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Customer Not Found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// GET /customers
func (h *CustomerHandler) GetAllCustomers(c *gin.Context) {

	customers, err := h.service.GetAllCustomers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, customers)
}

func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Customer ID",
		})
		return
	}

	var req db.UpdateCustomerParams

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	req.ID = int32(id)

	err = h.service.UpdateCustomer(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Customer Updated Successfully",
	})
}

func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Customer ID",
		})
		return
	}

	err = h.service.DeleteCustomer(c, int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Customer Deleted Successfully",
	})
}