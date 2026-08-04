package report

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetUsersAtCreditLimit(c *gin.Context) {
	users, err := h.service.GetUsersAtCreditLimit(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) GetCustomersWithDue(c *gin.Context) {
	customers, err := h.service.GetCustomersWithDue(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}

func (h *Handler) GetCustomerDueByName(c *gin.Context) {
	name := c.Param("name")

	due, err := h.service.GetCustomerDueByName(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, due)
}

func (h *Handler) GetAllMerchantsFeeCollected(c *gin.Context) {
	report, err := h.service.GetAllMerchantsFeeCollected(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, report)
}
