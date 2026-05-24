package post

import (
	"context"
	"time"

	"github.com/KelsonTeixeira/go/go-api/internal"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	Conn *pgxpool.Pool
}

func (r *Repository) Insert(post internal.Post) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.Conn.Exec(
		ctx,
		"INSERT INTO posts (username, body) VALUES ($1, $2)",
		post.Username,
		post.Body,
	)

	return err
}
