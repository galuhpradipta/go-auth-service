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
	router.Post("/api/user/login", handler.Login)
}

func (h handler) Register(ctx *fiber.Ctx) error {

	payload := models.UserRegisterRequest{}
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(models.HttpResponse{
			Error: err.Error(),
		})
	}

	user, err := h.userUsecase.Register(payload.Email, payload.Address, payload.Password)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.HttpResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(models.HttpResponse{
		Data: map[string]interface{}{
			"message": "successfully registering new user",
			"user_id": user.ID,
		},
	})
}

func (h handler) Login(ctx *fiber.Ctx) error {

	payload := models.UserLoginRequest{}
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(models.HttpResponse{
			Error: err.Error(),
		})
	}

	// user, err := h.userUsecase.Register(payload.Email, payload.Address, payload.Password)
	// if err != nil {
	// 	return ctx.Status(http.StatusInternalServerError).JSON(models.HttpResponse{
	// 		Error: err.Error(),
	// 	})
	// }

	return ctx.JSON(models.HttpResponse{
		Data: payload,
	})
}
