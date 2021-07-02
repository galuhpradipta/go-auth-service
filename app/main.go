package main

import (
	"fmt"
	"log"

	"github.com/galuhpradipta/go-auth-service/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// userRepository "github.com/galuhpradipta/go-auth-user/domain/user/repository"
)

func main() {
	fmt.Println("go auth service")

	db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})

	// fiber := fiber.New()

	// userRepository := userRepository.NewRepository()

}
