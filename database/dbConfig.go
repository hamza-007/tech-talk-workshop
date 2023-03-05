package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func Connect() {
	db, e := sql.Open("sqlite3", "./database/data.db")
	if e != nil {
		log.Fatalf("Error: %v", e)
	}

	if e := db.Ping(); e != nil {
		log.Fatalf("Error: %v", e)
	}
	DB = db
}