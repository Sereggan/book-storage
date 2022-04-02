package service

import (
	"context"
	"github.com/Sereggan/book-storage/internal/model"
	"github.com/Sereggan/book-storage/internal/store"
	"github.com/sirupsen/logrus"
)

type BookService struct {
	store store.Book
}

func NewBookService(store store.Book) *BookService {
	return &BookService{store}
}

func (b *BookService) GetByAuthorId(ctx context.Context, id uint64) ([]*model.Book, error) {
	books, err := b.store.FindByAuthor(ctx, id)
	if err != nil {
		logrus.Errorf("Failed to get books by authorId: %d, err: %v", id, err)
		return nil, err
	}

	return books, nil
}
