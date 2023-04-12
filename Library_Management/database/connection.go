package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@@tcp(localhost:3306)/chinni")
	if err != nil {
		panic(err.Error())
	}
	return db
}
