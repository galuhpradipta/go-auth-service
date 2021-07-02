package middleware

import (
	"github.com/galuhpradipta/go-auth-service/models"
	"github.com/galuhpradipta/go-auth-service/shared/jwt/constant"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(constant.JWT_SECRET_KEY),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(models.HttpResponse{Error: "Missing or malformed JWT"})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(models.HttpResponse{Error: "Invalid or expired JWT"})
}
