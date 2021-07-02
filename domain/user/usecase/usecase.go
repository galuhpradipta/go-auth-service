package usecase

import (
	"errors"

	"github.com/galuhpradipta/go-auth-service/domain/user"
	"github.com/galuhpradipta/go-auth-service/models"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailRecordExist = errors.New("email record already exist")
	ErrHashPassword     = errors.New("failed to hashpassword")
)

const BCRYPT_HASH_COST = 12

type usecase struct {
	userRepository user.Repository
}

func NewUsecase(userRepository user.Repository) user.Usecase {
	return &usecase{
		userRepository: userRepository,
	}
}

func (u usecase) Register(email, address, password string) (models.User, error) {

	user, _ := u.userRepository.FindByEmail(email)
	if user.ID != 0 {
		return models.User{}, ErrEmailRecordExist
	}

	hashedPassword, err := u.hashPassword(password)
	if err != nil {
		return models.User{}, ErrHashPassword
	}

	return u.userRepository.Create(models.User{
		Email:    email,
		Address:  address,
		Password: hashedPassword,
	})
}

func (u usecase) hashPassword(plain string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plain), BCRYPT_HASH_COST)
	if err != nil {
		return "", err

	}

	return string(hashedPassword), nil
}
