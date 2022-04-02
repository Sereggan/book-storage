package grpcserver

import (
	"context"
	"github.com/Sereggan/book-storage/internal/delivery/grpc/pb"
	"github.com/Sereggan/book-storage/internal/service"
)

type Server struct {
	author service.Author
	book   service.Book
	pb.BookStorageServer
}

func New(services *service.Service) *Server {
	return &Server{
		author: services.Author,
		book:   services.Book,
	}
}

func (s *Server) GetBooksByAuthorId(ctx context.Context, request *pb.BookRequest) (*pb.BookList, error) {
	response := &pb.BookList{}
	books, err := s.book.GetByAuthorId(ctx, request.AuthorId)
	if err != nil {
		response.Error = err.Error()
		return response, nil
	}
	response.Book = mapBookListToProtoBookList(books)

	return response, nil
}

func (s *Server) GetAuthorsByBookId(ctx context.Context, request *pb.AuthorRequest) (*pb.AuthorList, error) {
	response := &pb.AuthorList{}
	authors, err := s.author.GetByBookId(ctx, request.BookId)
	if err != nil {
		response.Error = err.Error()
		return response, nil
	}
	response.Author = mapAuthorListToProtoAuthorList(authors)

	return response, nil
}
