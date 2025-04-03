package service

import (
	"context"
	"database/sql"

	"github.com/redblood-pixel/FilesExchanger/internal/domain"
	"github.com/redblood-pixel/FilesExchanger/internal/repository"
)

type Files interface {
	UploadFile(ctx context.Context, file domain.File) (int, error)
	ListFiles(ctx context.Context) ([]domain.File, error)
	DownloadFile(ctx context.Context, filename string) ([]byte, error)
}

type Service struct {
	Files
}

func NewService(path string, repo *repository.Repository, db *sql.DB) *Service {
	return &Service{
		Files: NewFileService(path, repo.File, db),
	}
}
