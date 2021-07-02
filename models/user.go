package models

import "github.com/jinzhu/gorm"

type (
	User struct {
		gorm.Model
		Email    string
		Address  string
		Password string
	}

	UserRegisterRequest struct {
		Email    string
		Address  string
		Password string
	}
)
