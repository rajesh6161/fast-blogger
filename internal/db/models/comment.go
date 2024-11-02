package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID          uuid.UUID `json:"id"`
	Content     string    `json:"content"`
	ByID        uuid.UUID `gorm:"type:uuid" json:"by_id"`
	By          *User     `gorm:"foreignKey:ByID;constraint:OnDelete:CASCADE;" json:"by"`
	PostID      uuid.UUID `gorm:"type:uuid" json:"post_id"`
	Post        *Post     `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE;" json:"post"`
	DateCreated time.Time `json:"date_created"`
}
