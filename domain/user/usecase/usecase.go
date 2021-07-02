package usecase

import (
	"errors"

	"github.com/galuhpradipta/go-auth-service/domain/user"
	"github.com/galuhpradipta/go-auth-service/models"
	"golang.org/x/crypto/bcrypt"

	sharedJwt "github.com/galuhpradipta/go-auth-service/shared/jwt"
)

var (
	ErrEmailRecordExist     = errors.New("error, email record already exist")
	ErrHashPassword         = errors.New("error, failed to hashpassword")
	ErrPasswordDoesNotMatch = errors.New("error, password doesnt match")
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

func (u usecase) Login(email, plainPassword string) (string, error) {

	user, err := u.userRepository.FindByEmail(email)
	if err != nil {
		return "", err
	}

	if err := u.compareHashAndPassword(user.Password, plainPassword); err != nil {
		return "", ErrPasswordDoesNotMatch
	}

	token, err := sharedJwt.GenerateSessionToken(user.Email)
	if err != nil {
		return "", err
	}

	return token, nil

}

func (u usecase) hashPassword(plain string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plain), BCRYPT_HASH_COST)
	if err != nil {
		return "", err

	}

	return string(hashedPassword), nil
}

func (u usecase) compareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
