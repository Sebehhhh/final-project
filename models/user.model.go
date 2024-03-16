package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
    ID        uuid.UUID `gorm:"primaryKey"`
    Username  string
    Email     string
    Password  string
    Age       int
    CreatedAt time.Time
    UpdatedAt time.Time
    Photos    []Photo     
    Comments  []Comment   
    SocialMedias []SocialMedia 
}

type SignUpInput struct {
	Username        string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
	Age             int    `json:"age" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Age       int       `json:"age,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
