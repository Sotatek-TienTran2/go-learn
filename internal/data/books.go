package data

import "time"

type Book struct {
	ID          int64     `json:"id,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Author      string    `json:"author,omitempty"`
}
