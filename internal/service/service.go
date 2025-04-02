package service

import (
	"context"

	"github.com/redblood-pixel/FilesExchanger/internal/domain"
)

type Files interface {
	UploadFile(ctx context.Context, file domain.File) (int, error)
	ListFiles(ctx context.Context) ([]*domain.File, error)
	DownloadFile(ctx context.Context, filename string) ([]byte, error)
}

type Service struct {
	Files
}

func NewService(path string) *Service {
	return &Service{
		Files: NewFileService(path),
	}
}
