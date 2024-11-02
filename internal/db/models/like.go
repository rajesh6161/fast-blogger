package models

import "github.com/google/uuid"

type Like struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User   *User     `gorm:"foreignKey:UserID" json:"user"`
	PostID uuid.UUID `gorm:"type:uuid" json:"post_id"`
	Post   *Post     `gorm:"foreignKey:PostID" json:"post"`
}
