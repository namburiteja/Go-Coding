package main

import (
	"log"

	db "paylater-backend/internal/db"
	"paylater-backend/internal/database"
	"paylater-backend/internal/service"
	"paylater-backend/internal/handler"
	"paylater-backend/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	// Connect to MySQL
	conn, err := database.NewMySQLConnection() // this helps to conn in internal/database/mysql.go
	if err != nil {
		log.Fatal(err)
	}

	// Initialize sqlc
	queries := db.New(conn)

	customerService := service.NewCustomerService(queries)
	merchantService := service.NewMerchantService(queries)
	transactionService := service.NewTransactionService(queries)
	reportService := service.NewReportService(queries)


	// Handlers
	customerHandler := handler.NewCustomerHandler(customerService)
	merchantHandler := handler.NewMerchantHandler(merchantService)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	reportHandler := handler.NewReportHandler(reportService)

	// Gin
	router := gin.Default()

	// Routes
	routes.SetupRoutes(router, customerHandler, merchantHandler, transactionHandler, reportHandler)

	log.Println("Server Started on :9090")

	router.Run(":9090")

	log.Println("Database Connected Successfully")

	_ = queries
	
	_ = customerService
}