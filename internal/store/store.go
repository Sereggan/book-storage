package store

import (
	"book-storage/internal/model"
	"book-storage/internal/store/postgresql"
	"context"
	"github.com/jackc/pgx/v4"
)

type Book interface {
	FindByAuthor(ctx context.Context, authorId uint64) ([]*model.Book, error)
}

type Author interface {
	FindByBook(ctx context.Context, bookId uint64) ([]*model.Author, error)
}

type Store struct {
	Book
	Author
}

func NewStore(conn *pgx.Conn) *Store {
	return &Store{
		postgresql.NewBookStore(conn),
		postgresql.NewAuthorStore(conn),
	}
}
