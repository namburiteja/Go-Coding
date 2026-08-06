package main

import (
	"log"
	"path/filepath"

	"paylater/services/report/internal/report"
	"paylater/shared/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv(
		filepath.Join("services", "report", ".env"),
		".env",
		filepath.Join("..", ".env"),
	)

	customerServiceURL := config.RequireEnv("CUSTOMER_SERVICE_URL")
	merchantServiceURL := config.RequireEnv("MERCHANT_SERVICE_URL")
	ledgerServiceURL := config.RequireEnv("LEDGER_SERVICE_URL")

	reportService := report.NewService(
		report.NewCustomersAPI(customerServiceURL),
		report.NewMerchantsAPI(merchantServiceURL),
		report.NewLedgerAPI(ledgerServiceURL),
	)
	reportHandler := report.NewHandler(reportService)

	router := gin.Default()
	report.RegisterRoutes(router, reportHandler)

	addr := config.ListenAddr()
	log.Println("Report service started on", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
