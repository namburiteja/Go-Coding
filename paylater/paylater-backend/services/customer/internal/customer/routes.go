package customer

import (
	"paylater/shared/middleware"

	"github.com/gin-gonic/gin"
)

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
