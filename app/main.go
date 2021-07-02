package main

import (
	"fmt"
	"log"

	"github.com/galuhpradipta/go-auth-service/models"
	"github.com/gofiber/fiber/v2"

	userHandler "github.com/galuhpradipta/go-auth-service/domain/user/delivery/http"
	userRepository "github.com/galuhpradipta/go-auth-service/domain/user/repository"
	userUsecase "github.com/galuhpradipta/go-auth-service/domain/user/usecase"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("go auth service")
	fiber := fiber.New()

	db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})

	userRepository := userRepository.NewRepository(db)
	userUsecase := userUsecase.NewUsecase(userRepository)
	userHandler.NewHandler(fiber, userUsecase)

	log.Fatal(fiber.Listen(":3000"))
}
