package service

import (
	"context"
	"github.com/Sereggan/book-storage/internal/model"
	"github.com/Sereggan/book-storage/internal/store"
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
