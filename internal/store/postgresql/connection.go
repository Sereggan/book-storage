package postgresql

import (
	"context"
	"github.com/jackc/pgx/v4"
)

func GetConnection(ctx context.Context, databaseURL string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		return nil, err
	}
	if err = conn.Ping(ctx); err != nil {
		return nil, err
	}
	return conn, nil
}
