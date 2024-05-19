package models

import (
	"gofiber-marketplace/src/configs"

	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	UserID        uint   `json:"user_id" validate:"required"`
	User          User   `gorm:"foreignKey:UserID" validate:"-"`
	Name          string `json:"name" validate:"required,max=50"`
	MainAddress   string `json:"main_address" validate:"required"`
	DetailAddress string `json:"detail_address" validate:"required"`
	Phone         string `json:"phone" validate:"required,numeric,max=15"`
	PostalCode    string `json:"postal_code" validate:"required,numeric,max=8"`
	City          string `json:"city" validate:"required"`
}

func SelectAllAddresses() []*Address {
	var addresses []*Address
	configs.DB.Preload("User").Find(&addresses)
	return addresses
}

func SelectAddressbyId(id int) *Address {
	var address Address
	configs.DB.Preload("User").First(&address, "id = ?", id)
	return &address
}

func CreateAddress(address *Address) error {
	result := configs.DB.Create(&address)
	return result.Error
}

func UpdateAddress(id int, updatedAddress *Address) error {
	result := configs.DB.Model(&Address{}).Where("id = ?", id).Updates(updatedAddress)
	return result.Error
}

func DeleteAddress(id int) error {
	result := configs.DB.Delete(&Address{}, "id = ?", id)
	return result.Error
}
