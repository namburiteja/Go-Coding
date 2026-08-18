package connection

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func ConnectDB() (*sql.DB, error) {
	dsn := "go_user:go123@tcp(127.0.0.1:3306)/people_db?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("connected to the database successfully")
	return db, nil
}
