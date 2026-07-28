package routes

import (
	"paylater-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	customerHandler *handler.CustomerHandler,
) {

	// Customer Routes
	router.POST("/customers", customerHandler.CreateCustomer)

	router.GET("/customers/:id", customerHandler.GetCustomerByID)

	router.GET("/customers", customerHandler.GetAllCustomers)

	router.PUT("/customers/:id", customerHandler.UpdateCustomer)

	router.DELETE("/customers/:id", customerHandler.DeleteCustomer)
}