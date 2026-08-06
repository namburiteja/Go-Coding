package main

import (
	"log"
	"path/filepath"

	db "paylater/services/ledger/internal/db"
	"paylater/services/ledger/internal/ledger"
	"paylater/shared/config"
	"paylater/shared/database"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv(
		filepath.Join("services", "ledger", ".env"),
		".env",
		filepath.Join("..", ".env"),
	)

	conn, err := database.NewMySQLConnection()
	if err != nil {
		log.Fatal(err)
	}

	customerServiceURL := config.RequireEnv("CUSTOMER_SERVICE_URL")
	merchantServiceURL := config.RequireEnv("MERCHANT_SERVICE_URL")

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

	addr := config.ListenAddr()
	log.Println("Ledger service started on", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
