package main

import (
	"log"
	"path/filepath"

	"paylater/gateway/internal/routes"
	"paylater/shared/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// Own .env only (not repository root). Works from gateway/, gateway/cmd/, or repo root.
	config.LoadEnv(
		filepath.Join("gateway", ".env"),
		".env",
		filepath.Join("..", ".env"),
	)

	adminServiceURL := config.RequireEnv("ADMIN_SERVICE_URL")
	merchantServiceURL := config.RequireEnv("MERCHANT_SERVICE_URL")
	customerServiceURL := config.RequireEnv("CUSTOMER_SERVICE_URL")
	ledgerServiceURL := config.RequireEnv("LEDGER_SERVICE_URL")
	reportServiceURL := config.RequireEnv("REPORT_SERVICE_URL")
	addr := config.ListenAddr()

	router := gin.Default()
	routes.SetupRoutes(
		router,
		adminServiceURL,
		merchantServiceURL,
		customerServiceURL,
		ledgerServiceURL,
		reportServiceURL,
	)

	log.Println("Gateway started on", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
