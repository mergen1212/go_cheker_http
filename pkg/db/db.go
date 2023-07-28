package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func GetDB() (sql.DB, error) {
	db, err := sql.Open("sqlite3", "./storage.db")
	if err != nil {
		return *db, err
	}
	return *db, nil
}

func PrepareDB(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS
	user (
	    id integer PRIMARY KEY AUTOINCREMENT,
	    telegram_id integer NOT NULL
	);`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS
	project(
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    hash VARCHAR(32) NOT NULL,
	    user_id INTEGER NOT NULL,
	    updated INTEGER NULL,
        FOREIGN KEY(user_id) REFERENCES user(id)
	)`)

	if err != nil {
		return err
	}
	return nil
}
