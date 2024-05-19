package models

import (
	"gofiber-marketplace/src/configs"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=20"`
	Role     string `json:"role" validate:"oneof=seller customer"`
}

func SelectAllUsers() []*User {
	var users []*User
	configs.DB.Find(&users)
	return users
}

func SelectUserById(id int) *User {
	var user User
	configs.DB.First(&user, "id = ?", id)
	return &user
}

func SelectUserbyEmail(email string) *User {
	var user User
	configs.DB.First(&user, "email = ?", email)
	return &user
}

func CreateUser(user *User) (uint, error) {
	result := configs.DB.Create(&user)
	return user.ID, result.Error
}

func UpdateUser(id int, updatedUser *User) error {
	result := configs.DB.Model(&User{}).Where("id = ?", id).Updates(updatedUser)
	return result.Error
}

func DeleteUser(id int) error {
	result := configs.DB.Delete(&User{}, "id = ?", id)
	return result.Error
}
