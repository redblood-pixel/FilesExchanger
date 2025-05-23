//go:build db_init
// +build db_init

package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./file_storage.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTableSQL := `
    CREATE TABLE IF NOT EXISTS files (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        original_name TEXT NOT NULL UNIQUE,
        stored_name TEXT NOT NULL UNIQUE CHECK(LENGTH(stored_name) = 68 OR LENGTH(stored_name)=69),
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );
    CREATE INDEX IF NOT EXISTS idx_stored ON files(stored_name);
	CREATE INDEX IF NOT EXISTS idx_original ON files(original_name);
    `

	if _, err = db.Exec(createTableSQL); err != nil {
		log.Fatal("Failed to create table:", err)
	}
}
