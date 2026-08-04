package main

import (
	"log"

	"paylater/gateway/internal/routes"
	"paylater/shared/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	adminServiceURL := config.GetEnv("ADMIN_SERVICE_URL")
	if adminServiceURL == "" {
		adminServiceURL = "http://localhost:9091"
	}

	merchantServiceURL := config.GetEnv("MERCHANT_SERVICE_URL")
	if merchantServiceURL == "" {
		merchantServiceURL = "http://localhost:9092"
	}

	customerServiceURL := config.GetEnv("CUSTOMER_SERVICE_URL")
	if customerServiceURL == "" {
		customerServiceURL = "http://localhost:9093"
	}

	ledgerServiceURL := config.GetEnv("LEDGER_SERVICE_URL")
	if ledgerServiceURL == "" {
		ledgerServiceURL = "http://localhost:9094"
	}

	reportServiceURL := config.GetEnv("REPORT_SERVICE_URL")
	if reportServiceURL == "" {
		reportServiceURL = "http://localhost:9095"
	}

	router := gin.Default()
	routes.SetupRoutes(
		router,
		adminServiceURL,
		merchantServiceURL,
		customerServiceURL,
		ledgerServiceURL,
		reportServiceURL,
	)

	log.Println("Gateway started on :9090")
	if err := router.Run(":9090"); err != nil {
		log.Fatal(err)
	}
}
