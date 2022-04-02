package service

import (
	"context"
	"github.com/Sereggan/book-storage/internal/model"
	"github.com/Sereggan/book-storage/internal/store"
	"github.com/sirupsen/logrus"
)

type AuthorService struct {
	store store.Author
}

func NewAuthorService(store store.Author) *AuthorService {
	return &AuthorService{store}
}

func (a *AuthorService) GetByBookId(ctx context.Context, id uint64) ([]*model.Author, error) {
	authors, err := a.store.FindByBook(ctx, id)
	if err != nil {
		logrus.Errorf("Failed to get authors by bookId: %d, err: %v", id, err)
		return nil, err
	}

	return authors, nil
}
