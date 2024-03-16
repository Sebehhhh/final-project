package models

import (
	"time"

	"github.com/google/uuid"
)

type SocialMedia struct {
	ID             uuid.UUID `gorm:"primaryKey"`
	Name           string
	SocialMediaURL string
	UserID         uuid.UUID // Foreign Key
	User           User      `gorm:"foreignKey:UserID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CreateSocialMediaRequest struct {
	Name           string    `json:"name,omitempty"`
	SocialMediaURL string    `json:"social_media_url,omitempty"`
	UserID         uuid.UUID `json:"user_id,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}

type UpdateSocialMediaRequest struct {
	Name           string    `json:"name,omitempty"`
	SocialMediaURL string    `json:"social_media_url,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}
