package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:Teja@7483@(localhost:3306)/smart")
	if err != nil {
		fmt.Println("Error in connecting the database:", err)
	}
	return db
}
