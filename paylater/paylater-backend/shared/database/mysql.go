package database

import (
	"database/sql"
	"fmt"

	"paylater/shared/config"

	_ "github.com/go-sql-driver/mysql"
)

// NewMySQLConnection opens MySQL using DB_* environment variables.
// All of DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, and DB_NAME must be set.
// Each business service must set its own DB_NAME (database-per-service).
func NewMySQLConnection() (*sql.DB, error) {
	host := config.RequireEnv("DB_HOST")
	port := config.RequireEnv("DB_PORT")
	user := config.RequireEnv("DB_USER")
	password := config.RequireEnv("DB_PASSWORD")
	name := config.RequireEnv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
