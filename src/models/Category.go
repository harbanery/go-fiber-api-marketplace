package models

import (
	"gofiber-marketplace/src/configs"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name     string    `json:"name" validate:"required,max=50"`
	Image    string    `json:"image" validate:"required"`
	Slug     string    `json:"slug" validate:"required,lowercase"`
	Products []Product `json:"products"`
}

func SelectAllCategories(keyword, sort string) []*Category {
	var categories []*Category
	keyword = "%" + keyword + "%"
	configs.DB.Preload("Products").Order(sort).Where("name ILIKE ?", keyword).Find(&categories)
	return categories
}

func SelectCategoryById(id int) *Category {
	var category Category
	configs.DB.Preload("Products").First(&category, "id = ?", id)
	return &category
}

func SelectCategoryBySlug(slug string) *Category {
	var category Category
	configs.DB.Preload("Products").First(&category, "slug = ?", slug)
	return &category
}

func CreateCategory(category *Category) error {
	result := configs.DB.Create(&category)
	return result.Error
}

func UpdateCategory(id int, updatedCategory *Category) error {
	result := configs.DB.Model(&Category{}).Where("id = ?", id).Updates(updatedCategory)
	return result.Error
}

func DeleteCategory(id int) error {
	result := configs.DB.Delete(&Category{}, "id = ?", id)
	return result.Error
}
