package models

import (
	"time"

	"github.com/go-playground/validator/v10"
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
	Name           string `json:"name" validate:"required"`
	SocialMediaURL string `json:"social_media_url" validate:"required,url"`
}

type UpdateSocialMediaRequest struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaURL string `json:"social_media_url" validate:"required,url"`
}

func ValidateCreateSocialMediaRequest(input CreateSocialMediaRequest) error {
	validate := validator.New()
	return validate.Struct(input)
}

func ValidateUpdateSocialMediaRequest(input UpdateSocialMediaRequest) error {
	validate := validator.New()
	return validate.Struct(input)
}
