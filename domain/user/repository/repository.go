package repository

import (
	"github.com/galuhpradipta/go-auth-service/domain/user"
	"github.com/galuhpradipta/go-auth-service/models"
	"gorm.io/gorm"
)

type repository struct {
	mysqlSess *gorm.DB
}

func NewRepository(mysqlSess *gorm.DB) user.Repository {
	return &repository{
		mysqlSess: mysqlSess,
	}
}

func (r repository) Create(user models.User) (models.User, error) {
	err := r.mysqlSess.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, err
}
