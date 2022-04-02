package main

import (
	"book-storage/internal/delivery/grpc/pb"
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
)

const address = "127.0.0.1:5300"

// Test client for grpc
func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println("error:", err)
	}

	defer conn.Close()

	client := pb.NewBookStorageClient(conn)

	req, err := client.GetBooksByAuthorId(context.Background(), &pb.BookRequest{
		AuthorId: 1,
	})

	logrus.Println(req)
}
