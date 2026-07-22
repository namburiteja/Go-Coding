package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"go-sqlc/generated"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// Data Source Name
	dsn := "go_user:go123@tcp(localhost:3306)/company"

	// Open database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verify connection
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ Database Connected")

	// Create sqlc Queries object
	queries := generated.New(db)

	// Create a base context
	ctx := context.Background()

	// Fetch all employees
	employees, err := queries.GetEmployees(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nEmployees:")
	for _, emp := range employees {
		fmt.Printf("%+v\n", emp)
	}
}