package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)
func insertEmployee(db *sql.DB) {

	query := `
	INSERT INTO employee(name, age, salary, team)
	VALUES (?, ?, ?, ?)
	`

	result, err := db.Exec(
		query,
		"Sai",
		21,
		50000,
		"Ai/Ml",
	)

	if err != nil {
		log.Fatal(err)
	}

	id, _ := result.LastInsertId()

	fmt.Println("Employee Inserted")
	fmt.Println("Inserted ID:", id)
}
func getEmployees(db *sql.DB) {

	rows, err := db.Query("SELECT * FROM employee")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("\nEmployees")

	for rows.Next() {

		var id int
		var name string
		var age int
		var salary float64
		var team string

		err := rows.Scan(
			&id,
			&name,
			&age,
			&salary,
			&team,
		)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(
			"ID:%d Name:%s Age:%d Salary:%.2f Team:%s\n",
			id,
			name,
			age,
			salary,
			team,
		)
	}
}
func updateEmployee(db *sql.DB, id int64) {

	query := `
	UPDATE employee
	SET salary=?, team=?
	WHERE id=?
	`

	result, err := db.Exec(
		query,
		90000,
		"Golang",
		id,
	)

	if err != nil {
		log.Fatal(err)
	}

	rows, _ := result.RowsAffected()

	fmt.Println("Rows Updated:", rows)
}
func deleteEmployee(db *sql.DB, id int64) {

	query := `
	DELETE FROM employee
	WHERE id=?
	`

	result, err := db.Exec(query, id)

	if err != nil {
		log.Fatal(err)
	}

	rows, _ := result.RowsAffected()

	fmt.Println("Rows Deleted:", rows)
}

func main() {

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

	fmt.Println("✅ Database Connected")

	// Uncomment ONE function at a time

	// insertEmployee(db)
	// getEmployees(db)
	// updateEmployee(db, 4)
	// deleteEmployee(db, 3)

}