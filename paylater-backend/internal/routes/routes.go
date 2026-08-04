package routes

import (
	"paylater-backend/internal/customer"
	"paylater-backend/internal/ledger"
	"paylater-backend/internal/merchant"
	"paylater-backend/internal/platform/middleware"
	"paylater-backend/internal/platform/proxy"
	"paylater-backend/internal/report"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	customerHandler *customer.Handler,
	merchantHandler *merchant.Handler,
	ledgerHandler *ledger.Handler,
	reportHandler *report.Handler,
	adminServiceURL string,
) {
	customers := router.Group("/customers")
	{
		customers.POST("/register", customerHandler.RegisterCustomer)
		customers.POST("/login", customerHandler.LoginCustomer)

		customers.GET(
			"",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			customerHandler.GetAllCustomers,
		)

		customers.POST(
			"/purchase",
			middleware.AuthMiddleware(),
			middleware.CustomerOnly(),
			ledgerHandler.Purchase,
		)

		customers.POST(
			"/payback",
			middleware.AuthMiddleware(),
			middleware.CustomerOnly(),
			ledgerHandler.Payback,
		)

		customers.GET(
			"/me",
			middleware.AuthMiddleware(),
			middleware.CustomerOnly(),
			customerHandler.GetMyProfile,
		)

		customers.PUT(
			"/me",
			middleware.AuthMiddleware(),
			middleware.CustomerOnly(),
			customerHandler.UpdateMyProfile,
		)

		customers.GET(
			"/me/transactions",
			middleware.AuthMiddleware(),
			middleware.CustomerOnly(),
			ledgerHandler.GetMyTransactions,
		)

		customers.GET(
			"/:id",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			customerHandler.GetCustomerByID,
		)

		customers.PUT(
			"/:id",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			customerHandler.UpdateCustomer,
		)

		customers.DELETE(
			"/:id",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			customerHandler.DeleteCustomer,
		)
	}

	merchants := router.Group("/merchants")
	{
		merchants.POST("/register", merchantHandler.RegisterMerchant)
		merchants.POST("/login", merchantHandler.LoginMerchant)

		merchants.GET(
			"",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			merchantHandler.GetAllMerchants,
		)

		merchants.GET(
			"/me",
			middleware.AuthMiddleware(),
			middleware.MerchantOnly(),
			merchantHandler.GetMyProfile,
		)

		merchants.PUT(
			"/me",
			middleware.AuthMiddleware(),
			middleware.MerchantOnly(),
			merchantHandler.UpdateMyProfile,
		)

		merchants.GET(
			"/me/transactions",
			middleware.AuthMiddleware(),
			middleware.MerchantOnly(),
			ledgerHandler.GetMerchantTransactions,
		)

		merchants.PUT(
			"/:id/commission",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			merchantHandler.UpdateMerchantCommission,
		)

		merchants.GET(
			"/:id",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			merchantHandler.GetMerchantByID,
		)

		merchants.PUT(
			"/:id",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			merchantHandler.UpdateMerchant,
		)

		merchants.DELETE(
			"/:id",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			merchantHandler.DeleteMerchant,
		)
	}

	transactions := router.Group("/transactions")
	{
		transactions.GET(
			"",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			ledgerHandler.GetAllTransactions,
		)
	}

	reports := router.Group("/reports")
	{
		reports.Use(
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
		)

		reports.GET("/credit-limit", reportHandler.GetUsersAtCreditLimit)
		reports.GET("/customers-due", reportHandler.GetCustomersWithDue)
		reports.GET("/customer-due/:name", reportHandler.GetCustomerDueByName)
		reports.GET("/merchant-fees", reportHandler.GetAllMerchantsFeeCollected)
	}

	adminProxy := proxy.Forward(adminServiceURL)
	router.Any("/admins", adminProxy)
	router.Any("/admins/*path", adminProxy)
}
