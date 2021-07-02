package jwt

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/dgrijalva/jwt-go"
	"github.com/galuhpradipta/go-auth-service/shared/constant"
)

func GenerateSessionToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = "1"
	claims["exp"] = time.Now().Add(time.Hour * 24 * 3) // 3 day expiration
	claims["email"] = email

	tokenString, err := token.SignedString([]byte(constant.JWT_SECRET_KEY))
	if err != nil {
		return "", err

	}
	return tokenString, nil
}

func ExtractToken(c *fiber.Ctx) (string, error) {
	token, err := verifyToken(c)
	if err != nil {
		return "", err
	}

	return token.Claims.(jwt.MapClaims)["email"].(string), err
}

func extractAuthHeader(c *fiber.Ctx) string {
	bearerToken := c.Get("Authorization")

	// Normally Authorization HTTP header.
	onlyToken := strings.Split(bearerToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractAuthHeader(c)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(constant.JWT_SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
