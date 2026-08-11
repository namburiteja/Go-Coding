package report

import (
	"errors"
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
	body, err := h.service.GetUsersAtCreditLimit(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", body)
}

func (h *Handler) GetCustomersWithDue(c *gin.Context) {
	body, err := h.service.GetCustomersWithDue(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", body)
}

func (h *Handler) GetCustomerDueByName(c *gin.Context) {
	name := c.Param("name")

	body, err := h.service.GetCustomerDueByName(c.Request.Context(), name)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", body)
}

func (h *Handler) GetAllMerchantsFeeCollected(c *gin.Context) {
	report, err := h.service.GetAllMerchantsFeeCollected(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, report)
}
