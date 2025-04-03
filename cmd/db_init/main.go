//go:build db_init
// +build db_init

package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// TODO create good db

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
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        original_name TEXT NOT NULL UNIQUE,
        stored_name TEXT NOT NULL UNIQUE CHECK(LENGTH(stored_name) = 64),
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );
    CREATE INDEX IF NOT EXISTS idx_stored ON files(stored_name);
	CREATE TRIGGER update_files_timestamp
	AFTER UPDATE ON files
	FOR EACH ROW
	BEGIN
		UPDATE files SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
	END;
    `

	if _, err = db.Exec(createTableSQL); err != nil {
		log.Fatal("Failed to create table:", err)
	}
}
