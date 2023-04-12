package routers

import (
	a "Donkey/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookId     int    `json:"BookId"`
	BookName   string `json:"BookName"`
	Author     string `json:"Author"`
	Price      int    `json:"price"`
	AuthorId   int    `json:"AuthorId"`
	AuthorName string `json:AuthorName"`
	Noof_Books int    `json:"No_of_books"`
}

func GetAllBooks(c *gin.Context) {
	db := a.ConnectToDB()
	books := []Book{}
	rows, err := db.Query("select a.AuthorId,b.BookId,b.BookName,b.Author from BookDetails b join Author a where a.AuthorName=b.Author")
	if err != nil {
		fmt.Println("error in query", err)
	}
	for rows.Next() {
		var booker Book
		err := rows.Scan(&booker.AuthorId, &booker.BookId, &booker.BookName, &booker.Author)
		if err != nil {
			fmt.Println("Error in Scanning", err)
		}
		books = append(books, booker)
	}
	c.JSON(http.StatusOK, books)
}
