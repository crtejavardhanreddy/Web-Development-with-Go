package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

type customer struct {
	Cid   int    `json:"Cid"`
	CName string `json:"CName"`
	Age   int    `json:"Age"`
}

func main() {
	r := gin.Default()
	r.GET("/customers", getAllCustomers)
	r.GET("/customers/:id", CustomerById)
	r.POST("/update", UpdateCustomerById)
	r.Run(":8089")
}
func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:Teja@7483@(localhost:3306)/customer")
	if err != nil {
		fmt.Println("Error in connecting the database...", err)
	}
	// defer db.Close()
	return db
}

func getAllCustomers(c *gin.Context) {
	db := Connect()
	rows, err := db.Query("SELECT * FROM new")
	if err != nil {
		fmt.Println("Unable to execute the query", err)
	}
	// defer rows.Close()
	customers := []customer{}
	for rows.Next() {
		var cus customer
		err := rows.Scan(&cus.Cid, &cus.CName, &cus.Age)
		if err != nil {
			fmt.Println("Unavle to scan the data...", err)
		}
		customers = append(customers, cus)
	}
	c.JSON(200, customers)
	fmt.Println(customers)
}

func CustomerById(c *gin.Context) {
	id := c.Param("id")
	var cus customer
	db := Connect()
	err := db.QueryRow("SELECT * FROM new WHERE Cid=?", id).Scan(&cus.Cid, &cus.CName, &cus.Age)
	fmt.Println(cus.Age)
	if err != nil {
		fmt.Println("Unable to execute the query", err)
	}
	c.JSON(http.StatusOK, cus)
}

func UpdateCustomerById(c *gin.Context) {
	db := Connect()
	// id := c.Param("id")
	var customer customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		fmt.Println("Unable to decode the input", err)
	}
	err = db.QueryRow("INSERT INTO new (Cid,CName,Age)values(?,?,?)", customer.Cid, customer.CName, customer.Age).Scan(&customer.Cid, &customer.CName, customer.Age)
	if err != nil {
		fmt.Println("Unable to insert the data", err)
	}
	c.JSON(200, customer)
}
