package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/redblood-pixel/FilesExchanger/internal/domain"
)

type FileRepo struct {
	db *sql.DB
}

func NewFileRepository(db *sql.DB) *FileRepo {
	return &FileRepo{db: db}
}

func (r *FileRepo) GetName(ctx context.Context, filename string) (string, error) {
	row := r.db.QueryRow("SELECT stored_name FROM files WHERE original_name = $1", filename)
	var hashName string
	if err := row.Scan(&hashName); err != nil {
		return "", err
	}
	return hashName, nil
}

func (r *FileRepo) GetAll(ctx context.Context) ([]domain.File, error) {
	rows, err := r.db.Query("SELECT original_name, created_at, updated_at FROM files")
	if err != nil {
		return nil, err
	}
	files := make([]domain.File, 0)
	for rows.Next() {
		f := domain.File{}
		if err := rows.Scan(&f.Name, &f.CreatedAt, &f.UpdatedAt); err != nil {
			return nil, err
		}
		files = append(files, f)
	}
	return files, nil
}

func (r *FileRepo) CreateOrUpdate(
	ctx context.Context,
	tx *sql.Tx,
	filename string,
	hashName string,
) error {

	var exists bool
	err := tx.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM files WHERE stored_name = ?)",
		hashName,
	).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		_, err = tx.ExecContext(
			ctx,
			"INSERT INTO files (original_name, stored_name) VALUES ($1, $2)",
			filename, hashName,
		)
	} else {
		_, err = tx.ExecContext(
			ctx,
			"UPDATE files SET updated_at = $1",
			time.Now(),
		)
	}
	return err
}
