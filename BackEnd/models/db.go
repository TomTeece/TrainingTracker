package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS Users (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement.Exec()
	statement, _ = db.Prepare("INSERT INTO Users (firstname, lastname) VALUES (?, ?)")
	statement.Exec("Tom", "Teece")
}
