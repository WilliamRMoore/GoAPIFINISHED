package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("COULD NOT CONNECT TO DB")
	}

	DB = db

	DB.SetMaxOpenConns(12)
	DB.SetMaxIdleConns(6)
	createTables()
}

func createTables() {

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL 
	)
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("COULD NOT CREATE USERS TABLE")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		location TEXT NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("COULD NOT CREATE EVENTS TABLE")
	}

	createRegistrationTables := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createRegistrationTables)

	if err != nil {
		panic("COULD NOT CREATE REGISTRATIONS TABLE")
	}
}