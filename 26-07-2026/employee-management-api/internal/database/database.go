package main
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
) 
func connect() (*sql.DB,error) {
	dsn := "go_user:go123@tcp(localhost:3306)/company?ParseTime=true"
	db,err := sql.Open("mysql",dsn)
	if err!=nil {
		return nil,err
	}
	err = db.Ping()
	if err!=nil {
		return nil,err
	}
	fmt.Println("Database created successfully")
	return db,nil
}