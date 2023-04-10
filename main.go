package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type custom struct {
	Cid   int    `json:"cid"`
	Cname string `json:"cName"`
	Age   string `json:"age"`
}

func ConnectToDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:1234@@tcp(localhost:3306)/chinni")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getCustomerById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	db, err := ConnectToDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()
	rows, err := db.Query("select * from customer where cid=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	var customer custom
	for rows.Next() {
		if err := rows.Scan(&customer.Cid, &customer.Cname, &customer.Age); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, customer)

}

func getCustomers(c *gin.Context) {
	db, err := ConnectToDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()
	rows, err := db.Query("select * from customer")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	customers := []custom{}

	for rows.Next() {
		var customer custom
		if err := rows.Scan(&customer.Cid, &customer.Cname, &customer.Age); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		customers = append(customers, customer)

	}
	c.JSON(http.StatusOK, customers)

}

func UpdateById(c *gin.Context) {
	db, err := ConnectToDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()
	var data custom

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.QueryRow("insert into customer(cid,cName,age) values(?,?,?)", data.Cid, data.Cname, data.Age).Scan(&data.Cid, &data.Cname, &data.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data updated successfully"})
}

func main() {
	r := gin.Default()
	r.GET("/customers", getCustomers)
	r.GET("/customer/:id", getCustomerById)
	r.POST("customeru", UpdateById)
	r.Run(":8081")
}
