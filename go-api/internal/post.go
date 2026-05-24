package internal

import "time"

type Post struct {
	Username  string
	Body      string
	CreatedAt time.Time
}
