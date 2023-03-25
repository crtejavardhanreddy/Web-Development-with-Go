package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	// fmt.Print("PSQL connection....")
	db, err := sql.Open("postgres", "dbname=custInfo user=postgres password=Teja@7483 host=localhost port=5432 sslmode=disable")
	if err != nil {
		fmt.Print("Unable to connect the psql...", err)
	}
	defer db.Close()
	return db
}
