package models

import "time"

type Post struct {
	Id         int       `json:"id" db:"id"`
	Title      string    `json:"title" db:"title"`
	Content    string    `json:"content" db:"content"`
	Author     string    `json:"author" db:"author"`
	Created_at time.Time `json:"created_at" db:"createdAt"`
	Updated_at time.Time `json:"updated_at" db:"updatedAt"`
}
