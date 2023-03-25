package userinfo

import (
	"fmt"
	a "v1/connection"
)

// var db *sql.DB

func GetDetails() {
	a.Connect()
	fmt.Println("The data is displayed..")
	// db.Query("CREATE DATABASE CUSTOMERS")
}
