package main

import (
	"fmt"
	a "v1/connection"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World!")
	a.Connect()
	r := gin.Default()
}
