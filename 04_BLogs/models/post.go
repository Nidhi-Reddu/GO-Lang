package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Headline string    `json:"headline" db:"headline"`
	Content  string    `json:"content" db:"content"`
	Author   string    `json:"author" db:"author"`
	Category string    `json:"category" db:"category"`
	UserID    uuid.NullUUID `json:"-" db:"user_id"`
	CreatedAt time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" db:"updated_at"`
}
