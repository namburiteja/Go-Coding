package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"

	"people-api/connection"
	"people-api/db/generated"
	"people-api/handler"
	"people-api/service"
)

func main() {

	// 1. Connect to MySQL
	dbs, err := connection.ConnectDB()
	if err != nil {
		log.Println("Error connecting to the database:", err)
		return
	}
	defer dbs.Close()

	// 2. Create SQLC Queries
	queries := generated.New(dbs)

	// 3. Create Service
	personService := service.NewPersonService(queries)

	// 4. Create Handler
	personHandler := handler.NewPersonHandler(personService)

	// 5. Register routes
	http.HandleFunc("/people", personHandler.GetAllPeople)

	http.HandleFunc("/people/cursor", personHandler.GetPeopleCursor)

	http.HandleFunc("/people/token", personHandler.GetPeopleToken)

	http.HandleFunc("/people/search", personHandler.SearchPeopleByName)

	http.HandleFunc("/people/", personHandler.GetPeopleByID)

	// PUT endpoint for updating a person
	http.HandleFunc("/people/update", personHandler.UpdatePerson)

	// 6. CORS
	handlerWithCORS := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:5173",
		},
		AllowedMethods: []string{
			"GET",
			"PUT",
			"OPTIONS",
		},
		AllowedHeaders: []string{
			"Content-Type",
		},
	}).Handler(http.DefaultServeMux)

	// 7. Start server
	log.Println("Server running on http://localhost:8090")

	err = http.ListenAndServe(":8090", handlerWithCORS)
	if err != nil {
		log.Fatal(err)
	}
}