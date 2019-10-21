package bookRepository

import (
	"database/sql"
	"github.com/ajrmzcs/books/models"
)

type BookRepository struct {}


func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) ([]models.Book, error) {
	rows, err := db.Query("SELECT * FROM books")

	if err != nil {
		return []models.Book{}, err
	}

	for rows.Next() {
		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.Year)
		books = append(books, book)
	}

	if err != nil {
		return []models.Book{}, err
	}

	return books, nil
}

func (b BookRepository) GetBook(db *sql.DB, book models.Book, id int) (models.Book, error) {
	row := db.QueryRow("SELECT * FROM books WHERE id=?", id)

	err := row.Scan(&book.Id, &book.Title, &book.Author, &book.Year)

	return book, err
}

func (b BookRepository) CreateBook(db *sql.DB, book models.Book) (int64, error) {
	res, err := db.Exec("INSERT INTO books (title, author, year) VALUES(?,?,?)",
		book.Title, book.Author, book.Year)

	if err != nil {
		return 0, err
	}

	bookId, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return bookId, nil
}

func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) (int64, error) {
	res, err := db.Exec("UPDATE books set title=?, author=?, year=? WHERE id=?",
		&book.Title, &book.Author, &book.Year, &book.Id)

	if err != nil {
		return 0, err
	}

	rU, err := res.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rU, nil
}

func (b BookRepository) DeleteBook(db *sql.DB, id int) (int64, error) {
	res, err := db.Exec("DELETE FROM books WHERE id=?", id)

	if err != nil {
		return 0, err
	}

	rD, err := res.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rD, nil
}