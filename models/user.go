package models

import "github.com/jinzhu/gorm"

type (
	User struct {
		gorm.Model
		Email    string `gorm:"index:idx_user_email,unique"`
		Address  string
		Password string
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
