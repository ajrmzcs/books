package main

import (
	"database/sql"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/ajrmzcs/books/driver"
	_ "reflect"
	"strconv"
	_ "strconv"

	//"fmt"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
)

var books []Book
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
	r:= mux.NewRouter()

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	fmt.Println("Server is running on port 8000")

	log.Fatal(http.ListenAndServe(":8000",r))
}

func getBooks(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query("SELECT * FROM books")
	logFatal(err)
	defer rows.Close()

	var b Book
	books = []Book{}

	for rows.Next() {
		err := rows.Scan(&b.Id, &b.Title, &b.Author, &b.Year)
		logFatal(err)
		books = append(books, b)
	}

	_ = json.NewEncoder(w).Encode(books)

}

func getBook(w http.ResponseWriter, r *http.Request) {
	var b Book
	params := mux.Vars(r)
	idP, _ := strconv.Atoi(params["id"])

	row := db.QueryRow("SELECT * FROM books WHERE id=?", idP)
	err := row.Scan(&b.Id, &b.Title, &b.Author, &b.Year)
	logFatal(err)

	_ = json.NewEncoder(w).Encode(b)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var b Book

	_ = json.NewDecoder(r.Body).Decode(&b)

	res, err := db.Exec("INSERT INTO books (title, author, year) VALUES(?,?,?)",
		b.Title, b.Author, b.Year)
	logFatal(err)

	id, err := res.LastInsertId()
	logFatal(err)

	_ = json.NewEncoder(w).Encode(id)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var b Book

	_ = json.NewDecoder(r.Body).Decode(&b)

	res, err := db.Exec("UPDATE books set title=?, author=?, year=? WHERE id=?",
		&b.Title, &b.Author, &b.Year, &b.Id)
	logFatal(err)

	rU, err := res.RowsAffected()
	logFatal(err)

	_ = json.NewEncoder(w).Encode(rU)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	res, err := db.Exec("DELETE FROM books WHERE id=?", id)
	logFatal(err)

	rD, err := res.RowsAffected()
	logFatal(err)

	_ = json.NewEncoder(w).Encode(rD)
}

