package main

import (
	"fmt"
	a "v1/connection"
	b "v1/webserver"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World!")
	a.Connect()
	r := gin.Default()
	r.GET("/restaurant", b.GetAllRestaurants)
	r.GET("/restaurant/:id", b.GetRestaurantById)
	r.GET("/waiter/:id", b.GetWaiter)
	r.GET("/restaurant_waiters/:id", b.GetAllWaiters)
	r.Run(":8089")
}
