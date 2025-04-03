package app

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	fsv1 "github.com/redblood-pixel/FilesExchanger/gen/v1"
	v1 "github.com/redblood-pixel/FilesExchanger/internal/delivery/v1"
	"github.com/redblood-pixel/FilesExchanger/internal/repository"
	"github.com/redblood-pixel/FilesExchanger/internal/service"
	"github.com/redblood-pixel/FilesExchanger/pkg/sqlite"
	"google.golang.org/grpc"
)

var path = "files/"

func Run() {

	db, err := sqlite.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db)
	svc := service.NewService(path, repo, db)
	handler := v1.NewFileGRPCHandler(svc)
	lis, err := net.Listen("tcp", ":5051")
	if err != nil {
		log.Fatal("listen error")
	}

	limiter := v1.NewRateLimiter()

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(limiter.LimiterInterceptor),
	)
	fsv1.RegisterFileServiceServer(grpcServer, handler)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal("serve error")
		}
	}()

	log.Println("Server started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Server stoped")
}
