package ledger

import (
	"paylater/shared/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes mounts ledger-owned public APIs (same paths as before extraction).
func RegisterRoutes(router *gin.Engine, h *Handler) {
	router.POST(
		"/customers/purchase",
		middleware.AuthMiddleware(),
		middleware.CustomerOnly(),
		h.Purchase,
	)
	router.POST(
		"/customers/payback",
		middleware.AuthMiddleware(),
		middleware.CustomerOnly(),
		h.Payback,
	)
	router.GET(
		"/customers/me/transactions",
		middleware.AuthMiddleware(),
		middleware.CustomerOnly(),
		h.GetMyTransactions,
	)
	router.GET(
		"/merchants/me/transactions",
		middleware.AuthMiddleware(),
		middleware.MerchantOnly(),
		h.GetMerchantTransactions,
	)
	router.GET(
		"/transactions",
		middleware.AuthMiddleware(),
		middleware.AdminOnly(),
		h.GetAllTransactions,
	)

	internal := router.Group("/internal/transactions", middleware.InternalServiceAuth())
	{
		internal.GET("/reports/merchant-fees", h.GetMerchantFeesInternal)
	}
}
