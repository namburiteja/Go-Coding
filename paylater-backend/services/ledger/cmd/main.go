package main

import (
	"log"

	db "paylater/services/ledger/internal/db"
	"paylater/services/ledger/internal/ledger"
	"paylater/shared/config"
	"paylater/shared/database"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	conn, err := database.NewMySQLConnection()
	if err != nil {
		log.Fatal(err)
	}

	customerServiceURL := config.GetEnv("CUSTOMER_SERVICE_URL")
	if customerServiceURL == "" {
		customerServiceURL = "http://localhost:9093"
	}

	merchantServiceURL := config.GetEnv("MERCHANT_SERVICE_URL")
	if merchantServiceURL == "" {
		merchantServiceURL = "http://localhost:9092"
	}

	queries := db.New(conn)
	ledgerService := ledger.NewService(
		conn,
		queries,
		ledger.NewCustomerCreditHTTP(customerServiceURL),
		ledger.NewMerchantCommissionHTTP(merchantServiceURL),
	)
	ledgerHandler := ledger.NewHandler(ledgerService)

	router := gin.Default()
	ledger.RegisterRoutes(router, ledgerHandler)

	addr := config.GetEnv("LEDGER_SERVICE_ADDR")
	if addr == "" {
		addr = ":9094"
	}

	log.Println("Ledger service started on", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
