package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/ajrmzcs/books/models"
	"github.com/ajrmzcs/books/repositories/book"
	"github.com/ajrmzcs/books/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Controller struct{}

var books []models.Book

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var e models.Error
		books = []models.Book{}
		bookRepo := bookRepository.BookRepository{}
		books, err := bookRepo.GetBooks(db, book, books)

		if err != nil {
			e.Message = "Server Error"
			utils.SendErrors(w, http.StatusInternalServerError, e)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, books)
	}
}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var e models.Error

		params := mux.Vars(r)

		books = []models.Book{}
		bookRepo := bookRepository.BookRepository{}

		id, _ := strconv.Atoi(params["id"])

		book, err := bookRepo.GetBook(db, book, id)

		if err!= nil {
			if err == sql.ErrNoRows {
				e.Message = "Not found"
				utils.SendErrors(w, http.StatusNotFound, e)
				return
			} else {
				e.Message = "Server Error"
				utils.SendErrors(w, http.StatusInternalServerError, e)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, book)

	}
}

func (c Controller) CreateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var e models.Error

		_ = json.NewDecoder(r.Body).Decode(&book)

		if book.Title == "" || book.Author == "" || book.Year == "" {
			e.Message = "Required fields are missing"
			utils.SendErrors(w, http.StatusBadRequest, e)
			return
		}

		bookRepo := bookRepository.BookRepository{}
		bookId, err := bookRepo.CreateBook(db, book)

		if err != nil {
			e.Message = "Server Error"
			utils.SendErrors(w, http.StatusInternalServerError, e)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, bookId)
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var e models.Error

		_ = json.NewDecoder(r.Body).Decode(&book)

		if book.Id == 0 ||book.Title == "" || book.Author == "" || book.Year == "" {
			e.Message = "All fields are Required"
			utils.SendErrors(w, http.StatusBadRequest, e)
			return
		}

		bookRepo := bookRepository.BookRepository{}
		rU, err := bookRepo.UpdateBook(db, book)

		if err != nil {
			e.Message = "Server Error"
			utils.SendErrors(w, http.StatusInternalServerError, e)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, rU)
	}
}

func (c Controller) DeleteBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Error
		params := mux.Vars(r)
		bookRepo := bookRepository.BookRepository{}
		id, _ := strconv.Atoi(params["id"])

		rD, err := bookRepo.DeleteBook(db, id)

		if err != nil {
			e.Message = "Server Error"
			utils.SendErrors(w, http.StatusInternalServerError, e)
			return
		}

		if rD == 0 {
			e.Message = "Not Found"
			utils.SendErrors(w, http.StatusNotFound, e) //404
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, rD)
	}
}


