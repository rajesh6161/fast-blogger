package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID          uuid.UUID `json:"id"`
	Content     string    `json:"content"`
	ByID        uuid.UUID `gorm:"type:uuid" json:"by_id"`
	By          *User     `json:"by"`
	PostID      uuid.UUID `gorm:"type:uuid" json:"post_id"`
	Post        *Post     `gorm:"foreignKey:PostID" json:"post"`
	DateCreated time.Time `json:"date_created"`
}
