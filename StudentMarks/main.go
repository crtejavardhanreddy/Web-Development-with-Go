package main

import (
	a "v1/BulkLoad"
	b "v1/MarksProcessor"
	c "v1/Webserver"

	_ "github.com/go-sql-driver/mysql"
)

// type Student struct {
// 	Sid    int    `json:"sid"`
// 	SName  string `json:"Sname"`
// 	SMarks int    `json:"SMarks"`
// }

func main() {
	a.OpenFile()
	a.LoadData()
	b.Marks()
	c.WebServer()
}
