package main

import (
	"log"

	"github.com/galuhpradipta/go-auth-service/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	userHandler "github.com/galuhpradipta/go-auth-service/domain/user/delivery/http"
	userRepository "github.com/galuhpradipta/go-auth-service/domain/user/repository"
	userUsecase "github.com/galuhpradipta/go-auth-service/domain/user/usecase"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fiber := fiber.New()
	fiber.Use(recover.New())
	fiber.Use(logger.New())

	sqliteDB, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqliteDB.AutoMigrate(&models.User{})

	userRepository := userRepository.NewRepository(sqliteDB)
	userUsecase := userUsecase.NewUsecase(userRepository)
	userHandler.NewHandler(fiber, userUsecase)

	log.Fatal(fiber.Listen(":3000"))
}
