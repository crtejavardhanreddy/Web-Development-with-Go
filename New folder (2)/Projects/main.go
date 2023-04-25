package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	cid   int    `json:"cid"`
	cName string `json:"cName"`
	cage  int    `json:"cage"`
}

func main() {
	http.HandleFunc("/get/", getById)
	http.ListenAndServe(":5001", nil)
}

func getById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/get/"):])
	fmt.Println(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	//database connection
	var db *sql.DB
	db, err = sql.Open("mysql", "root:Teja@7483@tcp(127.0.0.1:3306)/splitbill")
	if err != nil {
		log.Fatal("Unable to connect the database", err)
	}
	var users []User
	rows, err := db.Query("SELECT cid, cName,cage FROM customer WHERE id=?", id)
	if err != nil {
		log.Fatal("Unable to connect the database", err)
	}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.cid, &user.cName, &user.cage)
		if err != nil {
			fmt.Print("Unable to store the table", err)
		}
		users = append(users, user)
	}
	jsonData, err := json.Marshal(users)
	fmt.Println(jsonData)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	//return in json format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
