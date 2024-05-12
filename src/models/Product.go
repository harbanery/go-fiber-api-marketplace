package models

import (
	"gofiber-marketplace/src/configs"

	"gorm.io/gorm"
)

type ProductCondition string

const (
	New  ProductCondition = "new"
	Used ProductCondition = "used"
)

type Product struct {
	gorm.Model
	Name        string           `json:"name"`
	Price       float64          `json:"price"`
	Stock       int              `json:"stock"`
	URLImage    string           `json:"url_image"`
	Size        uint             `json:"size"`
	Color       string           `json:"color"`
	Rating      uint             `json:"rating" gorm:"default:0"`
	Description string           `json:"description"`
	Condition   ProductCondition `gorm:"type:product_condition" json:"condition"`
	CategoryID  uint             `json:"category_id"`
	Category    Category         `gorm:"foreignKey:CategoryID"`
	SellerID    uint             `json:"seller_id"`
	Seller      Seller           `gorm:"foreignKey:SellerID"`
}

func SelectAllProducts() []*Product {
	var products []*Product
	configs.DB.Preload("Category").Preload("Seller").Find(&products)
	return products
}

func SelectProductById(id int) *Product {
	var product Product
	configs.DB.Preload("Category").Preload("Seller").First(&product, "id = ?", id)
	return &product
}

func CreateProduct(product *Product) error {
	result := configs.DB.Create(&product)
	return result.Error
}

func UpdateProduct(id int, updatedProduct *Product) error {
	result := configs.DB.Model(&Product{}).Where("id = ?", id).Updates(updatedProduct)
	return result.Error
}

func DeleteProduct(id int) error {
	result := configs.DB.Delete(&Product{}, "id = ?", id)
	return result.Error
}
