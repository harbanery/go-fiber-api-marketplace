package helpers

import (
	"gofiber-marketplace/src/configs"
	"gofiber-marketplace/src/models"
	"log"
)

func Migration() {
	// configs.DB.AutoMigrate(&models.Product{})
	err := configs.DB.AutoMigrate(
		&models.User{},
		&models.Seller{},
		&models.Customer{},
		&models.Product{},
		&models.Category{},
		&models.Address{},
	)

	if err != nil {
		// Handle the error, e.g., log it or panic
		log.Fatalf("Failed to auto migrate: %v", err)
	}
}
