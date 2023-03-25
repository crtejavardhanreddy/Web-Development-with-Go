package main

import (
	"fmt"
	a "v1/connection"
	b "v1/userinfo"
)

func main() {
	fmt.Println("Connection initiated...")
	a.Connect()
	fmt.Println("connection established...")
	b.GetDetails()
}
