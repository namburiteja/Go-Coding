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

	// ---------------- Database Connection ----------------

	dsn := "go_user:go123@tcp(localhost:3306)/company"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ Database Connected Successfully")

	queries := generated.New(db)
	ctx := context.Background()

	// =====================================================
	// Uncomment ONE function at a time
	// =====================================================

	// createEmployee(ctx, queries)

	// getEmployees(ctx, queries)

	// getEmployee(ctx, queries, 4)

	updateEmployee(ctx, queries)

	// deleteEmployee(ctx, queries, 1)
}

// -------------------------------------------------------
// CREATE
// -------------------------------------------------------

func createEmployee(ctx context.Context, queries *generated.Queries) {

	_, err := queries.CreateEmployee(ctx, generated.CreateEmployeeParams{
		Name:   "Anji",
		Age:    22,
		Salary: "75000",
		Team:   "Backend",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ Employee Inserted Successfully")
}

// -------------------------------------------------------
// READ ALL
// -------------------------------------------------------

func getEmployees(ctx context.Context, queries *generated.Queries) {

	employees, err := queries.GetEmployees(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n========== Employees ==========")

	for _, emp := range employees {

		fmt.Printf(
			"ID:%d Name:%s Age:%d Salary:%s Team:%s\n",
			emp.ID,
			emp.Name,
			emp.Age,
			emp.Salary,
			emp.Team,
		)
	}
}

// -------------------------------------------------------
// READ ONE
// -------------------------------------------------------

func getEmployee(ctx context.Context, queries *generated.Queries, id int32) {

	employee, err := queries.GetEmployee(ctx, id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n========== Employee ==========")

	fmt.Printf(
		"ID:%d Name:%s Age:%d Salary:%s Team:%s\n",
		employee.ID,
		employee.Name,
		employee.Age,
		employee.Salary,
		employee.Team,
	)
}

// -------------------------------------------------------
// UPDATE
// -------------------------------------------------------

func updateEmployee(ctx context.Context, queries *generated.Queries) {

	err := queries.UpdateEmployee(ctx, generated.UpdateEmployeeParams{
		Salary: "100000",
		Team:   "Golang",
		ID:     3,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ Employee Updated Successfully")
}

// -------------------------------------------------------
// DELETE
// -------------------------------------------------------

func deleteEmployee(ctx context.Context, queries *generated.Queries, id int32) {

	err := queries.DeleteEmployee(ctx, id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ Employee Deleted Successfully")
}