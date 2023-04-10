package BulkLoad

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	Sid    int    `json:"sid"`
	SName  string `json:"Sname"`
	SMarks int    `json:"SMarks"`
}

var Students []Student

func OpenFile() {
	studentinfo, err := os.Open("BulkLoad/studentMarks.csv")
	fmt.Println("File Readed")
	if err != nil {
		fmt.Println("Unable to read the CSV file", err)
	}
	reader := csv.NewReader(studentinfo)
	temp, _ := reader.ReadAll()
	for _, data := range temp {
		a, _ := strconv.Atoi(data[0])
		b, _ := strconv.Atoi(data[2])
		Students = append(Students, Student{Sid: a, SName: data[1], SMarks: b})
	}
}
func LoadData() {
	db, err := sql.Open("mysql", "root:Teja@7483@tcp(127.0.0.1:3306)/student")
	if err != nil {
		fmt.Println("Unable to open the database...", err)
	}
	fmt.Println("Connection established...")
	// defer db.Close()
	db.Query("DROP TABLE student")
	db.Query("CREATE TABLE student(Sid int,SName varchar(200),SMarks int) ")
	// fmt.Println("Data inserting...", Students)
	// fmt.Printf("%T", Students)
	for _, data := range Students {
		db.Query("INSERT INTO student(Sid,SName,SMarks)values(?,?,?)", data.Sid, data.SName, data.SMarks)
	}
	db.Query("UPDATE student SET Sid=1 WHERE Sid=0")
	defer db.Close()
}
