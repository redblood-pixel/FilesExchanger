//go:build db_init
// +build db_init

package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Открываем соединение (файл БД создастся автоматически)
	db, err := sql.Open("sqlite3", "./file_storage.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создаём таблицу для хранения информации о файлах
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS files (
        id TEXT PRIMARY KEY,
        original_name TEXT NOT NULL,
        stored_name TEXT NOT NULL UNIQUE,
        user_id TEXT NOT NULL,
        upload_time DATETIME DEFAULT CURRENT_TIMESTAMP,
        size INTEGER NOT NULL
    );
    CREATE INDEX IF NOT EXISTS idx_user ON files(user_id);
    `

	if _, err = db.Exec(createTableSQL); err != nil {
		log.Fatal("Failed to create table:", err)
	}
}
