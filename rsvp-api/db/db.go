package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db") // Assign to global DB

	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables() // Call createTables to initialize DB schema
}

func createTables() {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);
	
	`
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		location TEXT,
		dateTime DATETIME NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	`

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY (event_id) REFERENCES events(id),
		FOREIGN KEY (user_id) REFERENCES users(id),
		UNIQUE(event_id, user_id)
	);
	`
	queries := []string{createUsersTable, createEventsTable, createRegistrationsTable}

	for _, query := range queries {
		_, err := DB.Exec(query)
		if err != nil {
			log.Fatal("Could not create table:", err)
		}
	}
}
