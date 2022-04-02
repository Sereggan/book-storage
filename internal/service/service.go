package service

import (
	"book-storage/internal/model"
	"book-storage/internal/store"
	"context"
)

type Author interface {
	GetByBookId(context.Context, uint64) ([]*model.Author, error)
}

type Book interface {
	GetByAuthorId(context.Context, uint64) ([]*model.Book, error)
}

type Service struct {
	Author
	Book
}

func NewService(stores *store.Store) *Service {
	return &Service{
		NewAuthorService(stores.Author),
		NewBookService(stores.Book),
	}
}
