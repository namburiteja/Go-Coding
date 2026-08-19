package main

import (
	"github.com/rs/cors"
	"log"
	"net/http"
	"people-api/connection"
	"people-api/db/generated"
	"people-api/handler"
	"people-api/service"
)

func main() {
	// 1. Connect to MySQL
	db, err := connection.ConnectDB()
	if err != nil {
		log.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	// 2. Create SQLC Queries
	queries := generated.New(db)

	// 3. Create Service
	personService := service.NewPersonService(queries)

	// 4. Create Handler
	personHandler := handler.NewPersonHandler(personService)

	// 5. Register routes
	http.HandleFunc("/people", personHandler.GetAllPeople)
	http.HandleFunc("/people/search", personHandler.SearchPeopleByName)
	http.HandleFunc("/people/", personHandler.GetPeopleByID)

	log.Println("Server running on http://localhost:8090")

	handlerWithCORS := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{"Content-Type"},
	}).Handler(http.DefaultServeMux)

	err = http.ListenAndServe(":8090", handlerWithCORS)

	if err != nil {
		log.Fatal(err)
	}
}
