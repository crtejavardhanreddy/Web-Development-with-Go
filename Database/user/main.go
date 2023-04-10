package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

type Restaurant struct {
	Rid     int    `json:"Restaurant_Id"`
	RName   string `json:"Restaurant_Name"`
	R_Ph_No int    `json:"Restaurant_PhoneNumber"`
}

var restaurants = make(map[string]Restaurant)

func main() {
	fmt.Println("Connecting the database...")
	db, err = sql.Open("mysql", "root:Teja@7483@tcp(127.0.0.1:3306)/splitbill")
	if err != nil {
		fmt.Println("Error in connecting the database", err)
	}
	defer db.Close()
	http.ListenAndServe(":9080", nil)
}

func addWaiter(w http.ResponseWriter, r *http.Request) {
	var rest Restaurant
	err := json.NewDecoder(r.Body).Decode(&rest)
	if err != nil {
		fmt.Println("Unable to decode the data...")
	}

}
