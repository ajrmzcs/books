package main

import (
	"database/sql"
	"fmt"
	"github.com/ajrmzcs/books/controllers"
	"github.com/ajrmzcs/books/driver"
	"github.com/ajrmzcs/books/models"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	_ "reflect"
)

var books []models.Book
var db *sql.DB

func init() {
	_ = gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db = driver.ConnectDB()
	controller := controllers.Controller{}

	r:= mux.NewRouter()

	r.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	r.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	r.HandleFunc("/books", controller.CreateBook(db)).Methods("POST")
	r.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	r.HandleFunc("/books/{id}", controller.DeleteBook(db)).Methods("DELETE")

	fmt.Println("Server is running on port 8000")

	log.Fatal(http.ListenAndServe(":8000",r))
}


