package postgresql

import (
	"context"
	"fmt"
	"github.com/Sereggan/book-storage/internal/model"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

type BookStore struct {
	conn *pgx.Conn
}

func NewBookStore(conn *pgx.Conn) *BookStore {
	return &BookStore{conn}
}

func (b *BookStore) FindByAuthor(ctx context.Context, authorId uint64) ([]*model.Book, error) {
	rows, err := b.conn.Query(ctx,
		"SELECT book_id, title, description, publication_date FROM book where "+
			"book_id IN (SELECT book_id from author_book where author_id = $1)",
		authorId)
	if err != nil {
		return nil, fmt.Errorf("could not find all books bu author_id, error: %s", err)
	}
	var books []*model.Book

	for rows.Next() {
		var bookId uint64
		var title string
		var description string
		var publicationDate pgtype.Date
		err = rows.Scan(&bookId, &title, &description, &publicationDate)
		if err != nil {
			return nil, fmt.Errorf("could not parse books, error: %s", err)
		}
		books = append(books, &model.Book{
			Id:              bookId,
			Title:           title,
			Description:     description,
			PublicationDate: publicationDate.Time.String(),
		})
	}

	return books, nil
}
