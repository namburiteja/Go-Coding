package customer

import (
	"database/sql"
	"net/http"
	"strconv"

	db "paylater/services/customer/internal/db"
	"paylater/shared/middleware"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// RegisterRoutes mounts public customer APIs and internal credit endpoints for Ledger.
func RegisterRoutes(router *gin.Engine, h *Handler) {
	customers := router.Group("/customers")
	{
		customers.POST("/register", h.RegisterCustomer)
		customers.POST("/login", h.LoginCustomer)

		customers.GET(
			"",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			h.GetAllCustomers,
		)

		customers.GET(
			"/me",
			middleware.AuthMiddleware(),
			middleware.CustomerOnly(),
			h.GetMyProfile,
		)

		customers.PUT(
			"/me",
			middleware.AuthMiddleware(),
			middleware.CustomerOnly(),
			h.UpdateMyProfile,
		)

		customers.GET(
			"/:id",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			h.GetCustomerByID,
		)

		customers.PUT(
			"/:id",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			h.UpdateCustomer,
		)

		customers.DELETE(
			"/:id",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			h.DeleteCustomer,
		)
	}

	internal := router.Group("/internal/customers", middleware.InternalServiceAuth())
	{
		// Report routes before /:id so "reports" is not captured as an id.
		internal.GET("/reports/at-credit-limit", h.GetUsersAtCreditLimitInternal)
		internal.GET("/reports/with-due", h.GetCustomersWithDueInternal)
		internal.GET("/reports/due-by-name/:name", h.GetCustomerDueByNameInternal)

		internal.GET("/:id/credit", h.GetCreditInternal)
		internal.PUT("/:id/due", h.UpdateDueInternal)
		internal.PUT("/:id/block", h.BlockInternal)
	}
}

func (h *Handler) RegisterCustomer(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.RegisterCustomer(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Customer registered successfully"})
}

func (h *Handler) LoginCustomer(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.LoginCustomer(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: token})
}

func (h *Handler) GetMyProfile(c *gin.Context) {
	userID := c.MustGet("userID").(int32)

	customer, err := h.service.GetMyProfile(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (h *Handler) UpdateMyProfile(c *gin.Context) {
	userID := c.MustGet("userID").(int32)

	var req db.UpdateCustomerParams
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

func (h *Handler) GetCustomerByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Customer ID"})
		return
	}

	customer, err := h.service.GetCustomerByID(c.Request.Context(), int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Customer Not Found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (h *Handler) GetAllCustomers(c *gin.Context) {
	customers, err := h.service.GetAllCustomers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customers)
}

func (h *Handler) UpdateCustomer(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Customer ID"})
		return
	}

	var req db.UpdateCustomerParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ID = int32(id)

	err = h.service.UpdateCustomer(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer Updated Successfully"})
}

func (h *Handler) DeleteCustomer(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Customer ID"})
		return
	}

	err = h.service.DeleteCustomer(c.Request.Context(), int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer Deleted Successfully"})
}

func (h *Handler) GetCreditInternal(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Customer ID"})
		return
	}

	customer, err := h.service.GetCreditForUpdate(c.Request.Context(), int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, toCreditSnapshot(customer))
}

func (h *Handler) UpdateDueInternal(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Customer ID"})
		return
	}

	var req UpdateDueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateDue(c.Request.Context(), int32(id), req.TotalDue); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "due updated"})
}

func (h *Handler) BlockInternal(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Customer ID"})
		return
	}

	if err := h.service.BlockCustomer(c.Request.Context(), int32(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "customer blocked"})
}

func (h *Handler) GetUsersAtCreditLimitInternal(c *gin.Context) {
	users, err := h.service.GetUsersAtCreditLimit(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) GetCustomersWithDueInternal(c *gin.Context) {
	customers, err := h.service.GetCustomersWithDue(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}

func (h *Handler) GetCustomerDueByNameInternal(c *gin.Context) {
	name := c.Param("name")
	customer, err := h.service.GetCustomerDueByName(c.Request.Context(), name)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func toCreditSnapshot(customer db.Customer) CreditSnapshotResponse {
	resp := CreditSnapshotResponse{
		ID:             customer.ID,
		CreditLimit:    customer.CreditLimit,
		PaymentDueDate: customer.PaymentDueDate,
	}
	if customer.TotalDue.Valid {
		resp.TotalDue = &customer.TotalDue.String
	}
	if customer.Status.Valid {
		status := string(customer.Status.CustomersStatus)
		resp.Status = &status
	}
	return resp
}
