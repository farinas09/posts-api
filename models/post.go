package models

import "time"

type Post struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserId    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Posts []Post
