package webserver

import (
	"fmt"
	"net/http"
	"strconv"
	a "v1/connection"

	"github.com/gin-gonic/gin"
)

//	type restaurant struct {
//		Restaurant_id             int    `json:"restaurant_id"`
//		Restaurant_name           string `json:"restaurant_name"`
//		Restaurant_mobile         int64  `json:"restaurant_mobile"`
//		Restaurant_account_number string `json:"restaurant_account_number"`
//		Ifsc_code                 string `json:"ifsc_code"`
//		Bank_name                 string `json:"bank_name"`
//		Bank_address              string `json:"bank_address"`
//	}
type restaurant struct {
	Restaurant_id             int    `json:"restaurant_id"`
	Restaurant_name           string `json:"restaurant_name"`
	Restaurant_mobile         int64  `json:"restaurant_mobile"`
	Restaurant_account_number string `json:"restaurant_account_number"`
	Ifsc_code                 string `json:"ifsc_code"`
	Bank_name                 string `json:"bank_name"`
	Bank_address              string `json:"bank_address"`
	Waiter_id                 int    `json:"waiter_id"`
	Waiter_name               string `json:"waiter_name"`
	Waiter_mobile             int64  `json:"waiter_mobile"`
	Waiter_address            string `json:"waiter_address"`
}

type waiter struct {
	Waiter_id               int    `json:"waiter_id"`
	Waiter_name             string `json:"waiter_name"`
	Restaurant_id           int    `json:"restaurant_id"`
	Restaurant_name         string `json:"restaurant_name"`
	Restaurant_mobile       int64  `json:"restaurant_mobile"`
	Restaurant_bank_address string `json:"restaurant_bank_address"`
}

func GetAllRestaurants(c *gin.Context) {
	db := a.Connect()
	rows, err := db.Query("SELECT * FROM restaurant_details")
	if err != nil {
		fmt.Print("Error:Unable to execute the query", err)
	}
	restaurants := []restaurant{}
	for rows.Next() {
		var res restaurant
		err := rows.Scan(&res.Restaurant_id, &res.Restaurant_name, &res.Restaurant_mobile, &res.Restaurant_account_number, &res.Ifsc_code, &res.Bank_name, &res.Bank_address)
		if err != nil {
			fmt.Print("Error:Unable to scan the data", err)
		}
		restaurants = append(restaurants, res)
	}
	c.JSON(200, restaurants)
	defer db.Close()
}

func GetRestaurantById(c *gin.Context) {
	db := a.Connect()
	id := c.Param("id")
	var res restaurant
	err := db.QueryRow("Select * from restaurant_details where restaurant_id=?", id).Scan(&res.Restaurant_id, &res.Restaurant_name, &res.Restaurant_mobile, &res.Restaurant_account_number, &res.Ifsc_code, &res.Bank_name, &res.Bank_address)
	if err != nil {
		fmt.Println("Unable to do the query", err)
	}
	c.JSON(200, res)
}

//get restaurant_details from  waiter_id

func GetWaiter(c *gin.Context) {
	db := a.Connect()
	id := c.Param("id")
	var wai waiter
	err := db.QueryRow("select w.waiter_id,w.waiter_name,r.restaurant_id,r.restaurant_name,r.restaurant_mobile,r.bank_address from waiter w JOIN restaurant_details r on w.restaurant_id=r.restaurant_id where w.waiter_id=?", id).Scan(&wai.Waiter_id, &wai.Waiter_name, &wai.Restaurant_id, &wai.Restaurant_name, &wai.Restaurant_mobile, &wai.Restaurant_bank_address)
	if err != nil {
		fmt.Println("Unable to do the query", err)
	}
	c.JSON(200, wai)
}

// get all waiters in the given restaurant_id
func GetAllWaiters(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	db := a.Connect()
	waiters := []restaurant{}
	rows, err := db.Query("select r.restaurant_id,r.restaurant_name,r.restaurant_mobile,r.restaurant_account_number,r.Ifsc_code,r.Bank_name,r.Bank_address,w.Waiter_id,w.Waiter_name,w.Waiter_mobile,w.Waiter_address  from restaurant_details r join waiter w on r.restaurant_id=w.restaurant_id where w.restaurant_id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for rows.Next() {
		var res restaurant
		if err := rows.Scan(&res.Restaurant_id, &res.Restaurant_name, &res.Restaurant_mobile, &res.Restaurant_account_number, &res.Ifsc_code, &res.Bank_name, &res.Bank_address, &res.Waiter_id, &res.Waiter_name, &res.Waiter_mobile, &res.Waiter_address); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		waiters = append(waiters, res)
	}
	c.JSON(http.StatusOK, waiters)

}
