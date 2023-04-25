package routes

import (
	a "v1/pkg/controllers"

	"github.com/gorilla/mux"
)

func Route() {
	r := mux.NewRouter()
	r.HandleFunc("/book", a.CreateBook).Methods("POST")
	r.HandleFunc("/book", a.GetBook).Methods("GET")
	r.HandleFunc("/book/{bookId}", a.GetBookById).Methods("GET")
	r.HandleFunc("/book/{bookId}", a.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{bookId}", a.DeleteBook).Methods("DELETE")
}
