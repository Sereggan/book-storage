package grpcserver

import (
	"github.com/Sereggan/book-storage/internal/delivery/grpc/pb"
	"github.com/Sereggan/book-storage/internal/model"
)

// maps internal model to external
func mapBookListToProtoBookList(books []*model.Book) []*pb.Book {

	result := make([]*pb.Book, 0, len(books))
	for _, book := range books {
		result = append(result, &pb.Book{
			Id:              book.Id,
			Title:           book.Title,
			Description:     book.Description,
			PublicationDate: book.PublicationDate,
		})
	}
	return result
}

// maps internal model to external
func mapAuthorListToProtoAuthorList(authors []*model.Author) []*pb.Author {
	result := make([]*pb.Author, 0, len(authors))
	for _, author := range authors {
		result = append(result, &pb.Author{
			Id:      author.Id,
			Name:    author.Name,
			Surname: author.Surname,
		})
	}
	return result
}
