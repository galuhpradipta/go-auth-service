package http

import (
	"net/http"

	"github.com/galuhpradipta/go-auth-service/domain/user"
	"github.com/galuhpradipta/go-auth-service/models"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	userUsecase user.Usecase
}

func NewHandler(router *fiber.App, userUsecase user.Usecase) {
	handler := handler{
		userUsecase: userUsecase,
	}

	router.Post("/api/user/register", handler.Register)
}

func (h handler) Register(ctx *fiber.Ctx) error {

	payload := models.UserRegisterRequest{}
	err := ctx.BodyParser(&payload)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(models.HttpResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(models.HttpResponse{
		Data: map[string]string{
			"message": "successfully registering new user",
		},
	})
}
