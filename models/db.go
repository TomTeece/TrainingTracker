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

	statement1, _ := db.Prepare("CREATE TABLE IF NOT EXISTS Users (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement1.Exec()

	statement2, _ := db.Prepare("CREATE TABLE IF NOT EXISTS Climbing (id INTEGER PRIMARY KEY, userId INTEGER, date TEXT, activity TEXT)")
	statement2.Exec()
}
