package helpers

import (
	"gofiber-marketplace/src/configs"
	"gofiber-marketplace/src/models"
)

func Migration() {
	// configs.DB.AutoMigrate(&models.Product{})
	configs.DB.AutoMigrate(
		&models.User{},
		&models.Seller{},
		&models.Product{},
		&models.Category{},
	)
}
