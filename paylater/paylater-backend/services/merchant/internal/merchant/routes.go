package merchant

import (
	"paylater/shared/middleware"

	"github.com/gin-gonic/gin"
)

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

	// Service-to-service internal APIs (INTERNAL_SERVICE_TOKEN, not end-user JWT).
	internal := router.Group("/internal/merchants", middleware.InternalServiceAuth())
	{
		// Report route before /:id so "reports" is not captured as an id.
		internal.GET("/reports/names", h.GetMerchantNamesInternal)
		internal.GET("/:id/commission", h.GetCommissionInternal)
	}
}
