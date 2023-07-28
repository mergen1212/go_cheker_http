package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	id          int
	telegram_id int
}

type Project struct {
	id      int
	hash    string
	user_id int
}

var databasePath = "./storage.db"

func GetDB() (sql.DB, error) {
	db, err := sql.Open("sqlite3", databasePath)
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
