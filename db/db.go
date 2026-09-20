package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "feed.db")

	if err != nil {
		panic("Could not open database")
	}

	DB.SetMaxOpenConns(2)
	DB.SetMaxIdleConns(1)

	createTables()
}

func createTables() {
	createFeedsTableStatement := `
	CREATE TABLE IF NOT EXISTS feeds (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createFeedsTableStatement)
	if err != nil {
		panic("Failed to create feeds table: \n" + err.Error())
	}

	createPostsTableStatement := `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT NOT NULL,
		created_at DATETIME NOT NULL,
		feed_id INTEGER NOT NULL,
		FOREIGN KEY(feed_id) REFERENCES feeds(id)
	)
	`

	_, err = DB.Exec(createPostsTableStatement)
	if err != nil {
		panic("Failed to create users table: \n" + err.Error())
	}
}
