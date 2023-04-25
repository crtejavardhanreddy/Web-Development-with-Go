package main

import (
	"database/sql"
	"fmt"
	a "v1/Connection"

	"github.com/gin-gonic/gin"
)

type User struct {
	cid   int    //`json:"cid"`
	cname string //`json:"cname"`
	cage  int    //`json:"cage"`
}

var db *sql.DB

var users []User

func main() {
	fmt.Println("Main file execution...")
	a.Connect()
	r := gin.Default()

	if err := r.Run(":5000"); err != nil {
		panic(err.Error())
	}
	// rows, err := db.Query("SELECT * FROM customer")
	// fmt.Println("Teja")
	// if err != nil {
	// 	fmt.Println("Unable to fetch the data from database", err)
	// }
	// for rows.Next() {
	// 	var user User
	// 	err = rows.Scan(&user.cid, &user.cname, &user.cage)
	// 	if err != nil {
	// 		fmt.Println("Unable to scan the data", err)
	// 	}
	// 	users = append(users, user)
	// }
	// fmt.Print(users)
}
