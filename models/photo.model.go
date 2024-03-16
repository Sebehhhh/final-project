package models

import (
	"time"

	"github.com/google/uuid"
)

type Photo struct {
    ID        uuid.UUID `gorm:"primaryKey"`
    Title     string
    Caption   string
    PhotoURL  string
    UserID    uuid.UUID // Foreign Key
    User      User `gorm:"foreignKey:UserID"`
    CreatedAt time.Time
    UpdatedAt time.Time
    Comments  []Comment 
}

type CreatePhotoRequest struct {
	Title     string    `json:"title" binding:"required"`
	Caption   string    `json:"caption" binding:"required"`
	PhotoURL  string    `json:"photo_url" binding:"required"`
	UserID    uuid.UUID `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UpdatePhoto struct {
	Title     string    `json:"title,omitempty"`
	Caption   string    `json:"caption,omitempty"`
	PhotoURL  string    `json:"photo_url,omitempty"`
	UserID    uuid.UUID `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
