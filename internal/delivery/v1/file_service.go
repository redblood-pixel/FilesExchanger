package v1

import (
	"context"

	fsv1 "github.com/redblood-pixel/FilesExchanger/gen/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FileServiceServer struct {
	fsv1.UnimplementedFileServiceServer
	// +service
}

// TODO implement server
func (s *FileServiceServer) DownloadFile(
	ctx context.Context,
	in *fsv1.DownloadFileRequest,
) (*fsv1.DownloadFileResponse, error) {
	return nil, nil
}

func (s *FileServiceServer) ListFiles(
	ctx context.Context,
	_ emptypb.Empty,
) (*fsv1.ListFilesResponse, error) {
	return nil, nil
}

func (s *FileServiceServer) UploadFile(
	ctx context.Context,
	in *fsv1.UploadFileRequest,
) (*fsv1.UploadFileResponse, error) {
	return nil, nil
}