package main

import (
	"log"
	"path/filepath"

	db "paylater/services/report/internal/db"
	"paylater/services/report/internal/report"
	"paylater/shared/config"
	"paylater/shared/database"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv(
		filepath.Join("services", "report", ".env"),
		".env",
		filepath.Join("..", ".env"),
	)

	conn, err := database.NewMySQLConnection()
	if err != nil {
		log.Fatal(err)
	}

	queries := db.New(conn)
	reportService := report.NewService(queries)
	reportHandler := report.NewHandler(reportService)

	router := gin.Default()
	report.RegisterRoutes(router, reportHandler)

	addr := config.ListenAddr()
	log.Println("Report service started on", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
