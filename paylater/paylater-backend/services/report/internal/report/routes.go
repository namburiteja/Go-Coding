package report

import (
	"paylater/shared/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes mounts report APIs with the same public paths as before.
func RegisterRoutes(router *gin.Engine, h *Handler) {
	reports := router.Group("/reports")
	{
		reports.Use(
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
		)

		reports.GET("/credit-limit", h.GetUsersAtCreditLimit)
		reports.GET("/customers-due", h.GetCustomersWithDue)
		reports.GET("/customer-due/:name", h.GetCustomerDueByName)
		reports.GET("/merchant-fees", h.GetAllMerchantsFeeCollected)
	}
}
