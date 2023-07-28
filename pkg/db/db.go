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

type Database struct {
	d *sql.DB
}

var databasePath = "./storage.db"

func GetDB() (Database, error) {
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		return Database{db}, err
	}
	return Database{db}, nil
}

func (db *Database) exec(query string, args ...any) (sql.Result, error) {
	return db.d.Exec(query, args...)
}

func (db *Database) oneRow(query string, args ...any) *sql.Row {
	return db.d.QueryRow(query, args...)
}

func (db *Database) PrepareDB() error {
	_, err := db.exec(`
	CREATE TABLE IF NOT EXISTS
	user (
	    id integer PRIMARY KEY AUTOINCREMENT,
	    telegram_id integer NOT NULL
	);`)
	if err != nil {
		return err
	}
	_, err = db.exec(`
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

func (db *Database) GetUser(id int) (User, error) {
	var user User
	row := db.oneRow("select id,telegram_id from user where id = ?", id)
	if err := row.Scan(&user.id, &user.telegram_id); err != nil {
		return user, err
	}
	return user, nil
}
