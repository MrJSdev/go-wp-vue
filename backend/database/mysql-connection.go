package database

import (
	"database/sql"
	"fmt"

	// ignoring the mysql
	_ "github.com/go-sql-driver/mysql"
)

// GetDB to get the db connect with config
func GetDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbName := "extracatchy_wp"
	dbUser := "root"
	dbPassword := "root"
	dbHost := "127.0.0.1"
	dbPort := ":8889"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPassword+"@tcp("+dbHost+dbPort+")/"+dbName)
	if err != nil {
		fmt.Println("having problem to connnect mysql db", err)
	}
	// defer DB.Close()
	return
}
