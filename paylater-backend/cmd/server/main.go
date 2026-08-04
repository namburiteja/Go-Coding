package main

import (
	"log"

	"paylater-backend/internal/customer"
	db "paylater-backend/internal/db"
	"paylater-backend/internal/database"
	"paylater-backend/internal/ledger"
	"paylater-backend/internal/merchant"
	"paylater-backend/internal/platform/config"
	"paylater-backend/internal/report"
	"paylater-backend/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	conn, err := database.NewMySQLConnection()
	if err != nil {
		log.Fatal(err)
	}

	queries := db.New(conn)

	customerService := customer.NewService(queries)
	merchantService := merchant.NewService(queries)
	ledgerService := ledger.NewService(
		conn,
		queries,
		ledger.NewCustomerCreditSQLC,
		ledger.NewMerchantCommissionSQLC,
	)
	reportService := report.NewService(queries)

	customerHandler := customer.NewHandler(customerService)
	merchantHandler := merchant.NewHandler(merchantService)
	ledgerHandler := ledger.NewHandler(ledgerService)
	reportHandler := report.NewHandler(reportService)

	adminServiceURL := config.GetEnv("ADMIN_SERVICE_URL")
	if adminServiceURL == "" {
		adminServiceURL = "http://localhost:9091"
	}

	router := gin.Default()
	routes.SetupRoutes(
		router,
		customerHandler,
		merchantHandler,
		ledgerHandler,
		reportHandler,
		adminServiceURL,
	)

	log.Println("Server Started on :9090")
	if err := router.Run(":9090"); err != nil {
		log.Fatal(err)
	}
}
