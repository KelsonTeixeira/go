package database

import (
	"context"
	"time"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

var Conn *pgx.ConnPool

func NewConnection(connectionString string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Conn, err := pgxpool.Connect(ctx, connectionString)
	if err != nil {
		return nil, err
	}

	return Conn, nil
}
