package main

import (
	"log"
	"path/filepath"

	db "paylater/services/merchant/internal/db"
	"paylater/services/merchant/internal/merchant"
	"paylater/shared/config"
	"paylater/shared/database"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv(
		filepath.Join("services", "merchant", ".env"),
		".env",
		filepath.Join("..", ".env"),
	)

	conn, err := database.NewMySQLConnection()
	if err != nil {
		log.Fatal(err)
	}

	queries := db.New(conn)
	merchantService := merchant.NewService(queries)
	merchantHandler := merchant.NewHandler(merchantService)

	router := gin.Default()
	merchant.RegisterRoutes(router, merchantHandler)

	addr := config.ListenAddr()
	log.Println("Merchant service started on", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
