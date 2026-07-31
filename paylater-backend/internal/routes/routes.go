package routes

import (
	"paylater-backend/internal/handler"
	"paylater-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	customerHandler *handler.CustomerHandler,
	merchantHandler *handler.MerchantHandler,
	transactionHandler *handler.TransactionHandler,
	reportHandler *handler.ReportHandler,
	adminHandler *handler.AdminHandler,
) {

	// ===========================
	// Customer Routes
	// ===========================
	customers := router.Group("/customers")
	{
		// Public
		customers.POST("/register", customerHandler.RegisterCustomer)
		customers.POST("/login", customerHandler.LoginCustomer)

		// Admin
		customers.GET(
			"",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			customerHandler.GetAllCustomers,
		)

		// Customer
		customers.POST(
			"/purchase",
			middleware.AuthMiddleware(),
			middleware.CustomerOnly(),
			transactionHandler.Purchase,
		)

		customers.POST(
			"/payback",
			middleware.AuthMiddleware(),
			middleware.CustomerOnly(),
			transactionHandler.Payback,
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
			transactionHandler.GetMyTransactions,
		)

		// Admin (ID based)
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

	// ===========================
	// Merchant Routes
	// ===========================
	merchants := router.Group("/merchants")
	{
		// Public
		merchants.POST("/register", merchantHandler.RegisterMerchant)
		merchants.POST("/login", merchantHandler.LoginMerchant)

		// Admin
		merchants.GET(
			"",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			merchantHandler.GetAllMerchants,
		)

		// Merchant
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
			transactionHandler.GetMerchantTransactions,
		)

		// Admin (ID based)
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

	// ===========================
	// Transaction Routes
	// ===========================
	transactions := router.Group("/transactions")
	{
		transactions.GET(
			"",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			transactionHandler.GetAllTransactions,
		)
	}

	// ===========================
	// Report Routes
	// ===========================
	report := router.Group("/reports")
	{
		report.Use(
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
		)

		report.GET("/credit-limit", reportHandler.GetUsersAtCreditLimit)
		report.GET("/customers-due", reportHandler.GetCustomersWithDue)
		report.GET("/customer-due/:name", reportHandler.GetCustomerDueByName)
		report.GET("/merchant-fees", reportHandler.GetAllMerchantsFeeCollected)
	}

	// ===========================
	// Admin Routes
	// ===========================
	admins := router.Group("/admins")
	{
		// Public
		admins.POST("/register", adminHandler.RegisterAdmin)
		admins.POST("/login", adminHandler.LoginAdmin)

		// Admin
		admins.GET(
			"",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			adminHandler.GetAllAdmins,
		)

		admins.GET(
			"/:id",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			adminHandler.GetAdminByID,
		)

		admins.PUT(
			"/:id",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			adminHandler.UpdateAdmin,
		)

		admins.DELETE(
			"/:id",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			adminHandler.DeleteAdminByID,
		)
	}
}






// package routes

// import (
// 	"paylater-backend/internal/handler"
// 	"paylater-backend/internal/middleware"

// 	"github.com/gin-gonic/gin"
// )

// func SetupRoutes(
// 	router *gin.Engine,
// 	customerHandler *handler.CustomerHandler,
// 	merchantHandler *handler.MerchantHandler,
// 	transactionHandler *handler.TransactionHandler,
// 	reportHandler *handler.ReportHandler,
// 	adminHandler *handler.AdminHandler,
// ) {

// 	// Customer Routes
// 	customers := router.Group("/customers")
// 	{
// 		// Public
// 		customers.POST("/register", customerHandler.RegisterCustomer)
// 		customers.POST("/login", customerHandler.LoginCustomer)

// 		customers.GET(
// 			"",
// 			middleware.AuthMiddleware(),
// 			middleware.AdminOnly(),
// 			customerHandler.GetAllCustomers,
// 		)
// 		customers.POST(
// 			"/purchase",
// 			middleware.AuthMiddleware(),
// 			middleware.CustomerOnly(),
// 			transactionHandler.Purchase,
// 		)

// 		customers.POST(
// 			"/payback",
// 			middleware.AuthMiddleware(),
// 			middleware.CustomerOnly(),
// 			transactionHandler.Payback,
// 		)
// 		customers.GET(
// 			"/me/transactions",
// 			middleware.AuthMiddleware(),
// 			middleware.CustomerOnly(),
// 			transactionHandler.GetMyTransactions,
// 		)

// 		customers.GET(
// 			"/:id",
// 			middleware.AuthMiddleware(),
// 			middleware.AdminOnly(),
// 			customerHandler.GetCustomerByID,
// 		)

// 		customers.PUT(
// 			"/:id",
// 			middleware.AuthMiddleware(),
// 			middleware.AdminOnly(),
// 			customerHandler.UpdateCustomer,
// 		)

// 		customers.DELETE(
// 			"/:id",
// 			middleware.AuthMiddleware(),
// 			middleware.AdminOnly(),
// 			customerHandler.DeleteCustomer,
// 		)
// 		customers.GET(
// 			"/me",
// 			middleware.AuthMiddleware(),
// 			middleware.CustomerOnly(),
// 			customerHandler.GetMyProfile,
// 		)

// 		customers.PUT(
// 			"/me",
// 			middleware.AuthMiddleware(),
// 			middleware.CustomerOnly(),
// 			customerHandler.UpdateMyProfile,
// 		)
// 	}

// 	// Merchant Routes
// 	// router.PUT("/merchants/:id/commission", merchantHandler.UpdateMerchantCommission)

// 	merchants := router.Group("/merchants")
// 	{
// 		// Public
// 		merchants.POST("/register", merchantHandler.RegisterMerchant)
// 		merchants.POST("/login", merchantHandler.LoginMerchant)


// 		// Admin
// 		merchants.GET(
// 			"",
// 			middleware.AuthMiddleware(),
// 			middleware.AdminOnly(),
// 			merchantHandler.GetAllMerchants,
// 		)
// 		merchants.GET(
// 			"/me/transactions",
// 			middleware.AuthMiddleware(),
// 			middleware.MerchantOnly(),
// 			transactionHandler.GetMerchantTransactions,
// 		)
// 		merchants.GET(
// 			"/:id",
// 			middleware.AuthMiddleware(),
// 			middleware.AdminOnly(),
// 			merchantHandler.GetMerchantByID,
// 		)

// 		merchants.PUT(
// 			"/:id",
// 			middleware.AuthMiddleware(),
// 			middleware.AdminOnly(),
// 			merchantHandler.UpdateMerchant,
// 		)

// 		merchants.PUT(
// 			"/:id/commission",
// 			middleware.AuthMiddleware(),
// 			middleware.AdminOnly(),
// 			merchantHandler.UpdateMerchantCommission,
// 		)

// 		merchants.DELETE(
// 			"/:id",
// 			middleware.AuthMiddleware(),
// 			middleware.AdminOnly(),
// 			merchantHandler.DeleteMerchant,
// 		)

// 		merchants.GET(
// 			"/me",
// 			middleware.AuthMiddleware(),
// 			middleware.MerchantOnly(),
// 			merchantHandler.GetMyProfile,
// 		)

// 		merchants.PUT(
// 			"/me",
// 			middleware.AuthMiddleware(),
// 			middleware.MerchantOnly(),
// 			merchantHandler.UpdateMyProfile,
// 		)
	
// 	}

// 	transactions := router.Group("/transactions")
// 	{
// 		transactions.GET(
// 			"",
// 			middleware.AuthMiddleware(),
// 			middleware.AdminOnly(),
// 			transactionHandler.GetAllTransactions,
// 		)

// 	}


// 	report := router.Group("/reports")
// 	{
// 		report.GET("/credit-limit", reportHandler.GetUsersAtCreditLimit)
// 		report.GET("/customers-due", reportHandler.GetCustomersWithDue)
// 		report.GET("/customer-due/:name", reportHandler.GetCustomerDueByName)
// 		report.GET("/merchant-fees", reportHandler.GetAllMerchantsFeeCollected)
// 	}

// 	admins := router.Group("/admins")
// 	{
// 		admins.POST("/register", adminHandler.RegisterAdmin)
// 		admins.POST("/login", adminHandler.LoginAdmin)
		
// 		admins.GET(
// 			"",
// 			middleware.AuthMiddleware(),
// 			middleware.AdminOnly(),
// 			adminHandler.GetAllAdmins,
// 		)
// 		admins.GET(
// 			"/:id",
// 			middleware.AuthMiddleware(),
// 			middleware.AdminOnly(),
// 			adminHandler.GetAdminByID,
// 		)
// 		admins.PUT(
// 			"/:id",
// 			middleware.AuthMiddleware(),
// 			middleware.AdminOnly(),
// 			adminHandler.UpdateAdmin,
// 		)
// 		admins.DELETE(
// 			"/:id",
// 			middleware.AuthMiddleware(),
// 			middleware.AdminOnly(),
// 			adminHandler.DeleteAdminByID,
// 		)
// 	}

// }