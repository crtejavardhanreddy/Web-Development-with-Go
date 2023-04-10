package WebServer

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	B "v1/BulkLoad"

	_ "github.com/go-sql-driver/mysql"
)

type Result struct {
	Sid     int    `json:"Sid"`
	SResult string `json:"SResult"`
}

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:Teja@7483@tcp(127.0.0.1:3306)/student")
	if err != nil {
		fmt.Println("Unable to open the database...", err)
	}
	return db
}

func WebServer() {
	http.HandleFunc("/studentsInfo", GetAllStudents)
	http.HandleFunc("/student/", GetStudentById)
	http.HandleFunc("/result/", StudentResultById)
	fmt.Println("Server starting at 8099 port...")
	http.ListenAndServe(":8099", nil)
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	Students := []B.Student{}
	rows, err := db.Query("SELECT * FROM student")
	if err != nil {
		fmt.Println("Unable to Retrive...", err)
	}
	for rows.Next() {
		var S B.Student
		err = rows.Scan(&S.Sid, &S.SName, &S.SMarks)
		if err != nil {
			fmt.Println("Unable to Scan the data...", err)
		}
		Students = append(Students, S)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Students)
}

func GetStudentById(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Path[len("/student/"):]
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println("Unable to convert the id", err)
	}
	db := Connect()
	rows, err := db.Query("SELECT * FROM student where Sid=?", id)
	if err != nil {
		fmt.Println("Unable to student with id", err)
	}
	var S B.Student
	for rows.Next() {

		err = rows.Scan(&S.Sid, &S.SName, &S.SMarks)
		if err != nil {
			fmt.Println("Unable to Scan the data", err)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(S)
}
func StudentResultById(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Path[len("/result/"):]
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println("Error in reading the id", err)
	}
	db := Connect()
	rows, err := db.Query("SELECT * FROM studentResults where Sid=?", id)
	if err != nil {
		fmt.Println("Error in reading the data by id", err)
	}
	var S Result
	for rows.Next() {
		err = rows.Scan(&S.Sid, &S.SResult)
		if err != nil {
			fmt.Println("Error in scaning the data by id", err)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(S)
}
