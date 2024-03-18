package models

import (
	"time"
)

type User struct {
	ID              int64  `gorm:"primaryKey"`
	Username        string `gorm:"size:50;not null"`
	Email           string `gorm:"size:150;not null"`
	Password        string `gorm:"type:text;not null"`
	Age             int    `gorm:"not null"`
	ProfileImageURL string `gorm:"type:text"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Photos          []Photo
	Comments        []Comment
	SocialMedias    []SocialMedia
}

type SignUpInput struct {
	Username        string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	Age             int    `json:"age" binding:"required"`
	ProfileImageURL string `json:"profile_image_url"`
}

type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserResponse struct {
	ID              int64     `json:"id,omitempty"`
	Username        string    `json:"username,omitempty"`
	Email           string    `json:"email,omitempty"`
	Age             int       `json:"age,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	ProfileImageURL string    `json:"profile_image_url,omitempty"`
}
