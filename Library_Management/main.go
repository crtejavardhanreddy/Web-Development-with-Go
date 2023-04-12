package main

import (
	b "Donkey/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/bookers", b.GetAllBooks)
	r.Run(":8090")
}
