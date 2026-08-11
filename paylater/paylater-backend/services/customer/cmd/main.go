package main

import (
	"log"
	"path/filepath"

	"paylater/services/customer/internal/customer"
	db "paylater/services/customer/internal/db"
	"paylater/shared/config"
	"paylater/shared/database"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv(
		filepath.Join("services", "customer", ".env"),
		".env",
		filepath.Join("..", ".env"),
	)

	conn, err := database.NewMySQLConnection()
	if err != nil {
		log.Fatal(err)
	}

	queries := db.New(conn)
	customerService := customer.NewService(conn, queries)
	customerHandler := customer.NewHandler(customerService)

	router := gin.Default()
	customer.RegisterRoutes(router, customerHandler)

	addr := config.ListenAddr()
	log.Println("Customer service started on", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
