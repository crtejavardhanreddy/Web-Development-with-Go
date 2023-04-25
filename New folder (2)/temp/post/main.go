package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Restaurant struct {
	Rid     int    `json:"Rid"`
	RName   string `json:"RName"`
	R_Ph_No int    `json:"R_Ph_No"`
}

func main() {
	router := gin.Default()
	router.POST("/Restaurant", addRestaurant)
	router.Run(":8089")
}

func addRestaurant(context *gin.Context) {
	Restaurants := []Restaurant{}
	var res Restaurant
	var err error
	if err = context.BindJSON(&res); err != nil {
		return
	}
	Restaurants = append(Restaurants, res)
	context.IndentedJSON(http.StatusCreated, Restaurants)
	fmt.Println(Restaurants)
}
