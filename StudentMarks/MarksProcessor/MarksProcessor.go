package Marksprocessor

import (
	"database/sql"
	"fmt"
	S "v1/BulkLoad"

	_ "github.com/go-sql-driver/mysql"
)

func Marks() {
	db, err := sql.Open("mysql", "root:Teja@7483@(localhost:3306)/student")
	if err != nil {
		fmt.Println("Unable to open the data base While Processing the Marks", err)

	}
	db.Query("DROP TABLE studentresults")
	db.Query("CREATE TABLE StudentResults(Sid int,SResult varchar(200))")
	for _, data := range S.Students {
		if data.SMarks >= 70 {
			db.Query("INSERT INTO StudentResults(Sid,SResult)values(?,'Pass with Distinction')", data.Sid)
			fmt.Println(1)
		} else if data.SMarks < 70 && data.SMarks >= 40 {
			db.Query("INSERT INTO StudentResults(Sid,SResult)values(?,'Pass')", data.Sid)
		} else {
			db.Query("INSERT INTO StudentResults(Sid,SResult)values(?,'Better Luck Next Time')", data.Sid)
		}
	}
	db.Query("UPDATE studentresults SET Sid=1 WHERE Sid=0")
	defer db.Close()
}
