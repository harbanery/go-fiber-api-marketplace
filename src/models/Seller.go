package models

import (
	"gofiber-marketplace/src/configs"

	"gorm.io/gorm"
)

type Seller struct {
	gorm.Model
	UserID      uint      `json:"user_id"`
	User        User      `gorm:"foreignKey:UserID"`
	Name        string    `json:"name"`
	URLImage    string    `json:"url_image"`
	Phone       string    `json:"phone"`
	Description string    `json:"description"`
	Products    []Product `json:"products"`
}

func SelectAllSellers() []*Seller {
	var sellers []*Seller
	// configs.DB.Preload("User").Find(&sellers)
	configs.DB.Preload("User").Preload("Products").Find(&sellers)
	return sellers
}

func SelectSellerById(id int) *Seller {
	var seller Seller
	// configs.DB.Preload("User").First(&seller, "id = ?", id)
	configs.DB.Preload("User").Preload("Products").First(&seller, "id = ?", id)
	return &seller
}

func CreateSeller(seller *Seller) error {
	result := configs.DB.Create(&seller)
	return result.Error
}

func UpdateSeller(id int, updatedSeller *Seller) error {
	result := configs.DB.Model(&Seller{}).Where("id = ?", id).Updates(updatedSeller)
	return result.Error
}

func DeleteSeller(id int) error {
	result := configs.DB.Delete(&Seller{}, "id = ?", id)
	return result.Error
}
