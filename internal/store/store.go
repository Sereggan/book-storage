package store

import (
	"context"
	"github.com/Sereggan/book-storage/internal/model"
	"github.com/Sereggan/book-storage/internal/store/postgresql"
	"github.com/jackc/pgx/v4"
)

//go:generate mockgen -destination=./mock/store_mock.go -package=mock github.com/Sereggan/book-storage/internal/store Author,Book
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
