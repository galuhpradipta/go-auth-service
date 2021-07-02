package usecase

import "github.com/galuhpradipta/go-auth-service/domain/user"

type usecase struct {
	userRepository user.Repository
}

func NewUsecase(userRepository user.Repository) user.Usecase {
	return &usecase{
		userRepository: userRepository,
	}
}
