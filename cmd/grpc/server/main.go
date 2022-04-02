package main

import (
	"book-storage/internal/config"
	"book-storage/internal/delivery/grpc/pb"
	"book-storage/internal/server/grpcserver"
	"book-storage/internal/service"
	"book-storage/internal/store"
	"book-storage/internal/store/postgresql"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		logrus.Print("No .env file found")
	}
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	cfg := config.New()

	conn, err := postgresql.GetConnection(context.Background(), cfg.DbAddress)
	if err != nil {
		logrus.Fatal(err)
	}

	stores := store.NewStore(conn)
	services := service.NewService(stores)
	srv := grpcserver.New(services)

	listener, err := net.Listen("tcp", fmt.Sprint(":", cfg.GrpcPort))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterBookStorageServer(grpcServer, srv)

	go func() {
		err = grpcServer.Serve(listener)
		if err != nil {
			logrus.Fatalf("failed start server: %v", err)
		}
	}()
	stopChan := make(chan os.Signal)

	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	<-stopChan
	grpcServer.GracefulStop()
	err = conn.Close(context.Background())
	if err != nil {
		logrus.Fatal(err)
	}
}
