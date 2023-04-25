package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "root:Teja@7483@tcp(localhost)/simplerest?charse=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Error in connecting the database", err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
