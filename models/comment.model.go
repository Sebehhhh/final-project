package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	UserID    uuid.UUID // Foreign Key
	User      User      `gorm:"foreignKey:UserID"`
	PhotoID   uuid.UUID // Foreign Key
	Photo     Photo     `gorm:"foreignKey:PhotoID"`
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateCommentRequest struct {
	UserID    uuid.UUID `json:"user_id,omitempty"`
	PhotoID   uuid.UUID `json:"photo_id,omitempty"`
	Message   string    `json:"message,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UpdateCommentRequest struct {
	Message   string    `json:"message,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
