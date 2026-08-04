package main

import (
	"log"

	"paylater/services/merchant/internal/merchant"
	db "paylater/services/merchant/internal/db"
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

	queries := db.New(conn)
	merchantService := merchant.NewService(queries)
	merchantHandler := merchant.NewHandler(merchantService)

	router := gin.Default()
	merchant.RegisterRoutes(router, merchantHandler)

	addr := config.GetEnv("MERCHANT_SERVICE_ADDR")
	if addr == "" {
		addr = ":9092"
	}

	log.Println("Merchant service started on", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
