package models

import (
	"gofiber-marketplace/src/configs"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name     string    `json:"name"`
	URLImage string    `json:"url_image"`
	Slug     string    `json:"slug"`
	Products []Product `json:"products"`
}

func SelectAllCategories() []*Category {
	var categories []*Category
	configs.DB.Preload("Products").Find(&categories)
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
