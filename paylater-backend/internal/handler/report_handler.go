package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"paylater-backend/internal/service"
)

type ReportHandler struct {
	service *service.ReportService
}

func NewReportHandler(service *service.ReportService) *ReportHandler {
	return &ReportHandler{
		service: service,
	}
}

// Get Users At Credit Limit
func (h *ReportHandler) GetUsersAtCreditLimit(c *gin.Context) {

	users, err := h.service.GetUsersAtCreditLimit(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

// Get Customers With Due
func (h *ReportHandler) GetCustomersWithDue(c *gin.Context) {

	customers, err := h.service.GetCustomersWithDue(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, customers)
}

// Get Customer Due By Name
func (h *ReportHandler) GetCustomerDueByName(c *gin.Context) {

	name := c.Param("name")

	due, err := h.service.GetCustomerDueByName(c.Request.Context(), name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, due)
}

// Get All Merchants Fee Collected
func (h *ReportHandler) GetAllMerchantsFeeCollected(c *gin.Context) {

	report, err := h.service.GetAllMerchantsFeeCollected(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, report)
}