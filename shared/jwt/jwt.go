package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	SECRET_KEY = "secret"
)

func GenerateSessionToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = "1"
	claims["exp"] = time.Now().Add(time.Hour * 24 * 3) // 3 day expiration
	claims["email"] = email

	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err

	}
	return tokenString, nil
}
