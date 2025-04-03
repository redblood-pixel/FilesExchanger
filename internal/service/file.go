package service

import (
	"context"
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"unicode/utf8"

	"github.com/redblood-pixel/FilesExchanger/internal/domain"
	"github.com/redblood-pixel/FilesExchanger/internal/repository"
	"github.com/redblood-pixel/FilesExchanger/pkg/fileutil"
)

type FileService struct {
	path     string
	fileRepo repository.File
	db       *sql.DB
}

func NewFileService(path string, fileRepo repository.File, db *sql.DB) *FileService {
	return &FileService{path: path, fileRepo: fileRepo, db: db}
}

func (s *FileService) UploadFile(ctx context.Context, file domain.File) (int, error) {

	if err := fileutil.Validate(file.Name); err != nil {
		return 0, errFileBadName
	}
	hashName := fileutil.HashFilename(file.Name)

	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.IsolationLevel(0),
		ReadOnly:  false,
	})
	if err != nil {
		return 0, errInternalServer
	}
	defer tx.Rollback()

	log.Println(file.Name, hashName, utf8.RuneCountInString(hashName))
	err = s.fileRepo.CreateOrUpdate(ctx, tx, file.Name, hashName)
	if err != nil {
		log.Println("fileservice, createorupdate", err.Error())
		return 0, errInternalServer
	}

	newFile, err := os.Create(filepath.Join(s.path, hashName))
	if err != nil {
		return 0, errInternalServer
	}
	n, err := newFile.Write(file.Content)
	if err != nil {
		return 0, errInternalServer
	}
	if err := tx.Commit(); err != nil {
		return 0, errInternalServer
	}
	return n, nil
}

func (s *FileService) ListFiles(ctx context.Context) ([]domain.File, error) {
	return s.fileRepo.GetAll(ctx)
}

func (s *FileService) DownloadFile(ctx context.Context, filename string) ([]byte, error) {

	// for streaming we should use file.Read with buffer of size, for example, 4 Kb
	hashName, err := s.fileRepo.GetName(ctx, filename)
	if err != nil {
		log.Println("getname", err.Error())
		return nil, errFileNotFound
	}
	data, err := os.ReadFile(filepath.Join(s.path, hashName))
	if err != nil {
		log.Println("findName", err.Error())
		return nil, errFileNotFound
	}
	return data, nil
}
