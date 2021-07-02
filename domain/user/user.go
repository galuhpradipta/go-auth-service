package user

import "github.com/galuhpradipta/go-auth-service/models"

type Repository interface {
	Create(user models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
}

type Usecase interface {
	Register(email, address, password string) (models.User, error)
}
