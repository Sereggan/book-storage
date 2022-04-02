package main

import (
	"context"
	"github.com/Sereggan/book-storage/internal/delivery/grpc/pb"
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
	// check when there are records in books
	resBooks, err := client.GetBooksByAuthorId(context.Background(), &pb.BookRequest{
		AuthorId: 1,
	})
	logrus.Println(resBooks.Book)

	// check when there are no records in books
	resBooks, err = client.GetBooksByAuthorId(context.Background(), &pb.BookRequest{
		AuthorId: 123,
	})
	logrus.Println(len(resBooks.Book) == 0)

	// check when there are records in authors
	resAuthors, err := client.GetAuthorsByBookId(context.Background(), &pb.AuthorRequest{
		BookId: 1,
	})
	logrus.Println(resAuthors)

	// check when there are no records in authors
	resAuthors, err = client.GetAuthorsByBookId(context.Background(), &pb.AuthorRequest{
		BookId: 4,
	})
	logrus.Println(len(resAuthors.Author) == 0)
}
