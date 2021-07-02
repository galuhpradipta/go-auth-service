package repository

import (
	"github.com/galuhpradipta/go-auth-service/domain/user"
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
