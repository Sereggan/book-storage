package postgresql

import (
	"book-storage/internal/model"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

type AuthorStore struct {
	conn *pgx.Conn
}

func NewAuthorStore(conn *pgx.Conn) *AuthorStore {
	return &AuthorStore{conn}
}

func (b *AuthorStore) FindByBook(ctx context.Context, bookId uint64) ([]*model.Author, error) {
	rows, err := b.conn.Query(ctx,
		"SELECT author_id, name, surname FROM author where "+
			"author_id IN (SELECT author_id from author_book where book_id = $1)",
		bookId)
	if err != nil {
		return nil, fmt.Errorf("could not find all authors bu book_id, error: %s", err)
	}
	var authors []*model.Author

	for rows.Next() {
		var authorId uint64
		var name string
		var surname string
		err = rows.Scan(&authorId, &name, &surname)
		if err != nil {
			return nil, fmt.Errorf("could not parse authors, error: %s", err)
		}
		authors = append(authors, &model.Author{
			Id:      authorId,
			Name:    name,
			Surname: surname,
		})
	}

	return authors, nil
}
