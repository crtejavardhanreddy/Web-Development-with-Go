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
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"Age"`
}

type Restaurant struct {
	Rid     int    `json:"Rid"`
	RName   string `json:"RName"`
	R_Ph_No int    `json:"R_Ph_No"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:Teja@7483@tcp(127.0.0.1:3306)/user")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/user/", getUserByID)
	http.HandleFunc("/users", getAllUsers)
	http.HandleFunc("/restaurant", addRestaurant)

	log.Println("Server listening on port 8089...")
	log.Fatal(http.ListenAndServe(":8089", nil))
	fmt.Println("Database established...")
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/user/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT * FROM user WHERE ID=?", id)

	var user User
	err = row.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
	fmt.Printf("%+v", user)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	users := []User{}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	fmt.Printf("%+v", users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func addRestaurant(w http.ResponseWriter, r *http.Request) {
	var res Restaurant
	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		fmt.Println("Unable to decode the input", err)
	}
	insert, err := db.Prepare("Insert Into restaurant(Rid,RName,R_Ph_No)values(?,?,?)")
	if err != nil {
		fmt.Println("Unable to insert the data", err)
	}
	_, err = insert.Exec(res.Rid, res.RName, res.R_Ph_No)
	if err != nil {
		fmt.Println("Unable to scan the data", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data added successfully"))
}
