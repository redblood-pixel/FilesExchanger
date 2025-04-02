package service

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"unicode"

	"github.com/redblood-pixel/FilesExchanger/internal/domain"

	"github.com/djherbis/times"
)

type FileService struct {
	path string
}

func NewFileService(path string) *FileService {
	return &FileService{path: path}
}

var types = []string{".jpg", ".png", ".jpeg"}

func (s *FileService) UploadFile(ctx context.Context, file domain.File) (int, error) {
	if c := strings.ContainsFunc(file.Name, func(r rune) bool {
		return !unicode.IsDigit(r) && !unicode.IsLetter(r) && r != '.' && r != '_'
	}); c {
		return 0, errFileInvalidName
	}
	ext := filepath.Ext(file.Name)
	if pos := slices.Index(types, ext); pos == -1 {
		return 0, errFileInvalidType
	}

	newPath := filepath.Join(s.path, file.Name)
	newFile, err := os.Create(newPath)
	if err != nil {
		return 0, errFileExists
	}
	n, err := newFile.Write(file.Content)
	if err != nil {
		return 0, fmt.Errorf("error while writing %w", err)
	}
	return n, nil
}

func (s *FileService) ListFiles(ctx context.Context) ([]*domain.File, error) {
	files, err := os.ReadDir(s.path)
	if err != nil {
		return nil, err
	}

	res := make([]*domain.File, len(files))
	for i, file := range files {
		t, err := times.Stat(filepath.Join(s.path, file.Name()))
		if err != nil {
			return nil, err
		}
		res[i] = &domain.File{
			Name:      file.Name(),
			CreatedAt: t.BirthTime(),
			UpdatedAt: t.ModTime(),
		}
	}
	return res, nil
}

func (s *FileService) DownloadFile(ctx context.Context, filename string) ([]byte, error) {

	// for streaming we should use file.Read with buffer of size, for example, 4 Kb
	data, err := os.ReadFile(filepath.Join(s.path, filename))
	if err != nil {
		return nil, errFileNotFound
	}
	return data, nil
}
