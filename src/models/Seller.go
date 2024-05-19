package models

import (
	"gofiber-marketplace/src/configs"

	"gorm.io/gorm"
)

type Seller struct {
	gorm.Model
	UserID      uint      `json:"user_id" validate:"required"`
	User        User      `gorm:"foreignKey:UserID" validate:"-"`
	Name        string    `json:"name" validate:"required,max=50"`
	Image       string    `json:"image" validate:"required"`
	Phone       string    `json:"phone" validate:"required,numeric,max=15"`
	Description string    `json:"description" validate:"required"`
	Products    []Product `json:"products"`
}

func SelectAllSellers() []*Seller {
	var sellers []*Seller
	configs.DB.Preload("User").Preload("Products").Find(&sellers)
	return sellers
}

func SelectSellerById(id int) *Seller {
	var seller Seller
	configs.DB.Preload("User").Preload("Products").First(&seller, "id = ?", id)
	return &seller
}

func SelectSellerByUserId(id int) *Seller {
	var seller Seller
	configs.DB.Preload("User").Preload("Products").First(&seller, "user_id = ?", id)
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
