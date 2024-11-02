package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID          uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title       string     `gorm:"not null" json:"title" validate:"required"`
	Body        string     `gorm:"type:text;not null" json:"body" validate:"required"`
	Author      *User      `json:"author" validate:"required"`
	ImageUrl    string     `json:"image_url"`
	Likes       []*Like    `json:"likes"`
	Comments    []*Comment `json:"comments"`
	DateCreated time.Time  `gorm:"autoCreateTime" json:"date_created"`
	DateUpdated time.Time  `gorm:"autoUpdateTime" json:"date_updated"`
}
