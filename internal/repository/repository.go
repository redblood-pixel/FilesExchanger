package repository

import (
	"context"
	"database/sql"

	"github.com/redblood-pixel/FilesExchanger/internal/domain"
)

type File interface {
	GetName(ctx context.Context, filename string) (string, error)
	GetAll(ctx context.Context) ([]domain.File, error)
	CreateOrUpdate(ctx context.Context, tx *sql.Tx, filename, hashName string) error
}

type Repository struct {
	File
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		File: NewFileRepository(db),
	}
}
