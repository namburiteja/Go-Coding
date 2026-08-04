package main

import (
	"log"

	"paylater-backend/internal/admin"
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
	adminService := admin.NewService(queries)

	customerHandler := customer.NewHandler(customerService)
	merchantHandler := merchant.NewHandler(merchantService)
	ledgerHandler := ledger.NewHandler(ledgerService)
	reportHandler := report.NewHandler(reportService)
	adminHandler := admin.NewHandler(adminService)

	router := gin.Default()
	routes.SetupRoutes(
		router,
		customerHandler,
		merchantHandler,
		ledgerHandler,
		reportHandler,
		adminHandler,
	)

	log.Println("Server Started on :9090")
	if err := router.Run(":9090"); err != nil {
		log.Fatal(err)
	}
}
