package main

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"unique"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type CreatePostPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Storage struct {
	db *gorm.DB
}
