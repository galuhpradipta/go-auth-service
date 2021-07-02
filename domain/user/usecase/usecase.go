package usecase

import (
	"github.com/galuhpradipta/go-auth-service/domain/user"
	"github.com/galuhpradipta/go-auth-service/models"
)

type usecase struct {
	userRepository user.Repository
}

func NewUsecase(userRepository user.Repository) user.Usecase {
	return &usecase{
		userRepository: userRepository,
	}
}

func (u usecase) Register(email, address, password string) (models.User, error) {
	return u.userRepository.Create(models.User{
		Email:    email,
		Address:  address,
		Password: password,
	})
}
