package Connection

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	fmt.Println("Database Connection initiated...")
	db, err := sql.Open("mysql", "root:Teja@7483@tcp(127.0.0.1:3306)/splitBill")
	if err != nil {
		fmt.Print("Error in database connection", err)
	}
	fmt.Println("Database connected successfully...")
	DB = db
	// return DB
}
