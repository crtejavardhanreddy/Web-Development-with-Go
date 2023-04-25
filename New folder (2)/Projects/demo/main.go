package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func getUsers() ([]byte, error) {
	//database connection
	db, err := sql.Open("mysql", "root:Teja@7483@tcp(localhost:3306)/splitbill")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	//reading the table
	rows, err := db.Query("SELECT * FROM customer")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//storing in the structure
	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func main() {
	jsonData, err := getUsers()
	if err != nil {
		panic(err)
	}

	var users []User
	err = json.Unmarshal(jsonData, &users)
	if err != nil {
		panic(err)
	}

	// do something with the users
	fmt.Printf("%+v", users)
}
