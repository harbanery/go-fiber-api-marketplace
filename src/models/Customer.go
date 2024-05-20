package models

import (
	"gofiber-marketplace/src/configs"
	"time"

	"gorm.io/gorm"
)

type CustomerGender string

const (
	Male   CustomerGender = "male"
	Female CustomerGender = "female"
	Other  CustomerGender = "other"
)

type Customer struct {
	gorm.Model
	UserID      uint           `json:"user_id" validate:"required"`
	User        User           `gorm:"foreignKey:UserID" validate:"-"`
	Name        string         `json:"name" validate:"required,max=50"`
	Image       string         `json:"image" validate:"required"`
	Phone       string         `json:"phone" validate:"required,numeric,max=15"`
	Gender      CustomerGender `gorm:"type:customer_gender;default:other" json:"gender" validate:"required,oneof=male female other"`
	DateOfBirth time.Time      `json:"date_of_birth" validate:"required"`
}

func SelectAllCustomers() []*Customer {
	var customers []*Customer
	configs.DB.Preload("User").Find(&customers)
	return customers
}

func SelectCustomerById(id int) *Customer {
	var customer Customer
	configs.DB.Preload("User").First(&customer, "id = ?", id)
	return &customer
}

func SelectCustomerByUserId(id int) *Customer {
	var customer Customer
	configs.DB.Preload("User").First(&customer, "user_id = ?", id)
	return &customer
}

func CreateCustomer(customer *Customer) error {
	result := configs.DB.Create(&customer)
	return result.Error
}

func UpdateCustomer(id int, updatedCustomer *Customer) error {
	result := configs.DB.Model(&Customer{}).Where("id = ?", id).Updates(updatedCustomer)
	return result.Error
}

func DeleteCustomer(id int) error {
	result := configs.DB.Delete(&Customer{}, "id = ?", id)
	return result.Error
}
