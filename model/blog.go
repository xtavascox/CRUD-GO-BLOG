package model

import "github.com/google/uuid"

type Blog struct {
	ID    uuid.UUID `json:"id" gorm:"primaryKey"`
	Title string    `json:"title" gorm:"not null;column:title;size:255"`
	Post  string    `json:"post" gorm:"not null;column:post;size:255"`
}
