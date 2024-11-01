package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title" validate:"required"`
	Body        string    `json:"body" validate:"required"`
	Author      string    `json:"author" validate:"required"`
	ImageUrl    string    `json:"image_url"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

// type Post struct {
// 	ID          int    `json:"id"`
// 	Title       string `json:"title"`
// 	Body        string `json:"body"`
// 	Author      string `json:"author"`
// 	ImageUrl    string `json:"image_url"`
// 	DateCreated string `json:"date_created"`
// 	Likes       int    `json:"likes"`
// 	Dislikes    int    `json:"dislikes"`
// 	Comments []Comment `json:"comments"`
// 	DateUpdated string `json:"date_updated"`
// }
