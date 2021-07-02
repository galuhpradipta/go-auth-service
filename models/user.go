package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type (
	User struct {
		gorm.Model
		Email    string `gorm:"index:idx_user_email,unique"`
		Address  string
		Password string
	}

	UserDTO struct {
		ID        uint      `json:"id"`
		Email     string    `json:"email"`
		Address   string    `json:"address"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	UserRegisterRequest struct {
		Email    string
		Address  string
		Password string
	}

	UserLoginRequest struct {
		Email    string
		Password string
	}
)
