package models

import (
	"time"
)

type SocialMedia struct {
	ID             int64  `gorm:"primaryKey"`
	Name           string `gorm:"size:50;not null"`
	SocialMediaURL string `gorm:"type:text;not null"`
	UserID         int64  `gorm:"not null"`
	User           User   `gorm:"foreignKey:UserID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CreateSocialMediaRequest struct {
	Name           string    `json:"name,omitempty"`
	SocialMediaURL string    `json:"social_media_url,omitempty"`
	UserID         int64     `json:"user_id,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}

type UpdateSocialMediaRequest struct {
	Name           string    `json:"name,omitempty"`
	SocialMediaURL string    `json:"social_media_url,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}
