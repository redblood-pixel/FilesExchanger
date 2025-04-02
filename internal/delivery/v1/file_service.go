package v1

import (
	"context"
	"log"
	"time"

	fsv1 "github.com/redblood-pixel/FilesExchanger/gen/v1"
	"github.com/redblood-pixel/FilesExchanger/internal/domain"
	"github.com/redblood-pixel/FilesExchanger/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TODO better status codes & error handling
// TODO hash file name and search by it

type rateLimiter struct {
	updateDownloadLimit chan struct{}
	listFilesLimit      chan struct{}
}

func NewRateLimiter() *rateLimiter {
	return &rateLimiter{
		updateDownloadLimit: make(chan struct{}, 10),
		listFilesLimit:      make(chan struct{}, 100),
	}
}

func (r *rateLimiter) LimiterInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	var limitChan chan struct{}
	switch info.FullMethod {
	case "/github.redbloodpixel.filesexchange.fileservice.v1.FileService/ListFiles":
		limitChan = r.listFilesLimit
	default:
		limitChan = r.updateDownloadLimit
	}

	waitCtx, cancel := context.WithTimeout(ctx, 10*time.Second) // important - ctx as parent context
	defer cancel()

	select {
	case limitChan <- struct{}{}:
		defer func() { <-limitChan }()
		return handler(ctx, req)
	case <-waitCtx.Done():
		return nil, status.Error(codes.ResourceExhausted, "timeout reached")
	}
}

type FileGRPCHandler struct {
	fsv1.UnimplementedFileServiceServer
	svc *service.Service
}

func NewFileGRPCHandler(svc *service.Service) *FileGRPCHandler {
	return &FileGRPCHandler{
		svc: svc,
	}
}

func (s *FileGRPCHandler) DownloadFile(
	ctx context.Context,
	in *fsv1.DownloadFileRequest,
) (*fsv1.DownloadFileResponse, error) {
	data, err := s.svc.Files.DownloadFile(ctx, in.Filename)
	if err != nil {
		return &fsv1.DownloadFileResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &fsv1.DownloadFileResponse{Content: &fsv1.FileContent{Data: data}}, nil
}

func (s *FileGRPCHandler) ListFiles(
	ctx context.Context,
	_ *emptypb.Empty,
) (*fsv1.ListFilesResponse, error) {
	files, err := s.svc.Files.ListFiles(ctx)
	if err != nil {
		return &fsv1.ListFilesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	out := make([]*fsv1.FileMetadata, len(files))
	for i, file := range files {
		out[i] = &fsv1.FileMetadata{
			Filename:  file.Name,
			CreatedAt: timestamppb.New(file.CreatedAt),
			UpdatedAt: timestamppb.New(file.UpdatedAt),
		}
	}
	return &fsv1.ListFilesResponse{Files: out}, nil

}

func (s *FileGRPCHandler) UploadFile(
	ctx context.Context,
	in *fsv1.UploadFileRequest,
) (*fsv1.UploadFileResponse, error) {
	if in == nil {
		log.Println("empty in")
		return &fsv1.UploadFileResponse{}, status.Error(codes.InvalidArgument, "message is empty")
	} else if in.Content == nil || len(in.Content.Data) == 0 {
		log.Println("empty in.File")
		return &fsv1.UploadFileResponse{}, status.Error(codes.InvalidArgument, "file info or content is empty")
	}
	file := domain.File{
		Name:    in.Filename,
		Content: in.Content.Data,
	}
	n, err := s.svc.Files.UploadFile(ctx, file)
	if err != nil {
		return &fsv1.UploadFileResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &fsv1.UploadFileResponse{Size: int32(n)}, nil
}
