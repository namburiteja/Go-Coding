package routes

import (
	"paylater-backend/internal/handler"

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

	// Customer Routes
	router.POST("/customers", customerHandler.CreateCustomer)
	router.GET("/customers", customerHandler.GetAllCustomers)
	router.GET("/customers/:id", customerHandler.GetCustomerByID)
	router.PUT("/customers/:id", customerHandler.UpdateCustomer)
	router.DELETE("/customers/:id", customerHandler.DeleteCustomer)

	// Merchant Routes
	router.POST("/merchants", merchantHandler.CreateMerchant)
	router.GET("/merchants", merchantHandler.GetAllMerchants)
	router.GET("/merchants/:id", merchantHandler.GetMerchantByID)
	router.PUT("/merchants/:id/commission", merchantHandler.UpdateMerchantCommission)
	router.PUT("/merchants/:id", merchantHandler.UpdateMerchant)
	router.DELETE("/merchants/:id",merchantHandler.DeleteMerchant)

	//Transaction Routes
	router.POST("/transactions", transactionHandler.CreateTransaction)
	router.GET("/transactions", transactionHandler.GetAllTransactions)
	router.GET("/transactions/customer/:id", transactionHandler.GetTransactionsByCustomerID)
	router.GET("/transactions/merchant/:id", transactionHandler.GetTransactionsByMerchantID)

	report := router.Group("/reports")
	{
		report.GET("/credit-limit", reportHandler.GetUsersAtCreditLimit)
		report.GET("/customers-due", reportHandler.GetCustomersWithDue)
		report.GET("/customer-due/:name", reportHandler.GetCustomerDueByName)
		report.GET("/merchant-fees", reportHandler.GetAllMerchantsFeeCollected)
	}

	admins := router.Group("/admins")
	{
		admins.POST("", adminHandler.CreateAdmin)
		admins.GET("", adminHandler.GetAllAdmins)
		admins.GET("/:id", adminHandler.GetAdminByID)
		admins.PUT("/:id", adminHandler.UpdateAdmin)
		admins.DELETE("/:id", adminHandler.DeleteAdminByID)
	}

}