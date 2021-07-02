package repository

import (
	"github.com/galuhpradipta/go-auth-service/domain/user"
	"github.com/galuhpradipta/go-auth-service/models"
	"gorm.io/gorm"
)

type repository struct {
	sqliteSess *gorm.DB
}

func NewRepository(sqliteSess *gorm.DB) user.Repository {
	return &repository{
		sqliteSess: sqliteSess,
	}
}

func (r repository) Create(user models.User) (models.User, error) {
	err := r.sqliteSess.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, err
}

func (r repository) FindByEmail(email string) (models.User, error) {
	user := models.User{}
	err := r.sqliteSess.Table("users").Where("email = ?", email).Find(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
