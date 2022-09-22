package database

import (
	"database/sql"
	"fmt"
	"os"

	// ignoring the mysql
	_ "github.com/go-sql-driver/mysql"
)

// GetDB to get the db connect with config
func GetDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbName := os.Getenv("DATABASE_NAME")
	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := ":" + os.Getenv("DATABASE_PORT")

	db, err = sql.Open(dbDriver, dbUser+":"+dbPassword+"@tcp("+dbHost+dbPort+")/"+dbName)
	if err != nil {
		fmt.Println("having problem to connnect mysql db", err)
	}
	// defer DB.Close()
	return
}
