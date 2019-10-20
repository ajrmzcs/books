package driver

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql",
		os.Getenv("DB_USER")+":"+os.Getenv("DB_USER")+"@/"+os.Getenv("DB_NAME"))

	logFatal(err)

	return db
}