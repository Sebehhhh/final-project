package models

import (
	"time"
)

type Comment struct {
	ID        int64  `gorm:"primaryKey"`
	UserID    int64  `gorm:"not null"`
	User      User   `gorm:"foreignKey:UserID"`
	PhotoID   int64  `gorm:"not null"`
	Photo     Photo  `gorm:"foreignKey:PhotoID"`
	Message   string `gorm:"size:200;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateCommentRequest struct {
	UserID    int64     `json:"user_id,omitempty"`
	PhotoID   int64     `json:"photo_id,omitempty"`
	Message   string    `json:"message,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UpdateCommentRequest struct {
	Message   string    `json:"message,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
