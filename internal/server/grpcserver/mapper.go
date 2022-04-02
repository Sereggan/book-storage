package grpcserver

import (
	"book-storage/internal/delivery/grpc/pb"
	"book-storage/internal/model"
)

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
