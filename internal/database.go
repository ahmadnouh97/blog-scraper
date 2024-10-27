package internal

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db/test.db")
	if err != nil {
		log.Fatal("Failed to open database: ", err)
	}

	// Set dialect
	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Fatal("Failed to set dialect: ", err)
	}

	// Run migrations
	if err := goose.Up(db, "./db/migrations"); err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}

	return db, nil
}
