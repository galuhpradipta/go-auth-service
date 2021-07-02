package http

import (
	"github.com/galuhpradipta/go-auth-service/domain/user"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	userUsecase user.Usecase
}

func NewHandler(router fiber.Router, userUsecase user.Usecase) {
	handler := handler{
		userUsecase: userUsecase,
	}

	_ = handler

	router.Delete("/", handler.Test)

}

func (h handler) Test(ctx *fiber.Ctx) error {

	return ctx.JSON("hello")
}
