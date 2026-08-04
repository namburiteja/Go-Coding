package main

import (
	"log"

	"paylater/services/admin/internal/admin"
	db "paylater/services/admin/internal/db"
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
	adminService := admin.NewService(queries)
	adminHandler := admin.NewHandler(adminService)

	router := gin.Default()
	admin.RegisterRoutes(router, adminHandler)

	addr := config.GetEnv("ADMIN_SERVICE_ADDR")
	if addr == "" {
		addr = ":9091"
	}

	log.Println("Admin service started on", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
