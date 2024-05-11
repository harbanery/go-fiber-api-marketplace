package controllers

import (
	"fmt"
	"gofiber-marketplace/src/models"
	"reflect"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllProduct(c *fiber.Ctx) error {
	products := models.SelectAllProducts()
	if len(products) == 0 {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"status":     "no content",
			"statusCode": 202,
			"message":    "Product is empty. You should create product",
		})
	}

	resultProducts := make([]*map[string]interface{}, len(products))
	for i, product := range products {
		resultProducts[i] = &map[string]interface{}{
			"id":            product.ID,
			"created_at":    product.CreatedAt,
			"updated_at":    product.UpdatedAt,
			"category_id":   product.CategoryID,
			"category_name": product.Category.Name,
			"brand_id":      product.SellerID,
			"brand_name":    product.Seller.Name,
			"name":          product.Name,
			"photo":         product.URLImage,
			"rating":        product.Rating,
			"price":         product.Price,
			"stock":         product.Stock,
			"desc":          product.Description,
		}
	}

	// return c.JSON(products)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"statusCode": 200,
		"data":       resultProducts,
	})
}

func GetDetailProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid ID format",
		})
	}

	product := models.SelectProductById(id)
	if product.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Product not found",
		})
	}

	resultProduct := map[string]interface{}{
		"id":            product.ID,
		"created_at":    product.CreatedAt,
		"updated_at":    product.UpdatedAt,
		"category_id":   product.CategoryID,
		"category_name": product.Category.Name,
		"brand_id":      product.SellerID,
		"brand_name":    product.Seller.Name,
		"name":          product.Name,
		"photo":         product.URLImage,
		"rating":        product.Rating,
		"price":         product.Price,
		"stock":         product.Stock,
		"desc":          product.Description,
	}

	// return c.JSON(product)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"statusCode": 200,
		"data":       resultProduct,
	})
}

func CreateProduct(c *fiber.Ctx) error {
	var newProduct models.Product

	if err := c.BodyParser(&newProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid request body",
		})
	}

	if reflect.TypeOf(newProduct.Name).Kind() != reflect.String {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Product name must be string",
		})
	} else if newProduct.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Product name cannot be empty",
		})
	}

	if _, err := strconv.ParseFloat(fmt.Sprintf("%f", newProduct.Price), 64); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Product price must be float",
		})
	} else if newProduct.Price <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Product price must be greater than zero",
		})
	}

	if _, err := strconv.Atoi(fmt.Sprintf("%d", newProduct.Stock)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Product price must be integer",
		})
	} else if newProduct.Stock < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Product stock cannot be negative",
		})
	}

	if newProduct.CategoryID != 0 {
		category := models.SelectCategoryById(int(newProduct.CategoryID))
		if category.ID == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":     "not found",
				"statusCode": 404,
				"message":    "Category not found",
			})
		}
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Category cannot be empty",
		})
	}

	if newProduct.SellerID != 0 {
		seller := models.SelectSellerById(int(newProduct.SellerID))
		if seller.ID == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":     "not found",
				"statusCode": 404,
				"message":    "Seller not found",
			})
		}
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Seller cannot be empty",
		})
	}

	if err := models.CreateProduct(&newProduct); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    "Failed to create product",
		})
	} else {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"status":     "success",
			"statusCode": 200,
			"message":    "Product created successfully",
		})
	}
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid ID format",
		})
	}

	product := models.SelectProductById(id)
	if product.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Product not found",
		})
	}

	var updatedProduct models.Product

	if err := c.BodyParser(&updatedProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid request body",
		})
	}

	if reflect.TypeOf(updatedProduct.Name).Kind() != reflect.String {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Product name must be string",
		})
	} else if updatedProduct.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Product name cannot be empty",
		})
	}

	if _, err := strconv.ParseFloat(fmt.Sprintf("%f", updatedProduct.Price), 64); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Product price must be float",
		})
	} else if updatedProduct.Price <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Product price must be greater than zero",
		})
	}

	if _, err := strconv.Atoi(fmt.Sprintf("%d", updatedProduct.Stock)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Product stock must be integer",
		})
	} else if updatedProduct.Stock < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Product stock cannot be negative",
		})
	}

	if updatedProduct.CategoryID != 0 {
		category := models.SelectCategoryById(int(updatedProduct.CategoryID))
		if category.ID == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":     "not found",
				"statusCode": 404,
				"message":    "Category not found",
			})
		}
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Category cannot be empty",
		})
	}

	if updatedProduct.SellerID != 0 {
		seller := models.SelectSellerById(int(updatedProduct.SellerID))
		if seller.ID == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":     "not found",
				"statusCode": 404,
				"message":    "Seller not found",
			})
		}
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Seller cannot be empty",
		})
	}

	if err := models.UpdateProduct(id, &updatedProduct); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    fmt.Sprintf("Failed to update product with ID %d", id),
		})
	} else {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"status":     "success",
			"statusCode": 200,
			"message":    fmt.Sprintf("Product with ID %d updated successfully", id),
		})
	}
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid ID format",
		})
	}

	product := models.SelectProductById(id)
	if product.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Product not found",
		})
	}

	if err := models.DeleteProduct(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    fmt.Sprintf("Failed to delete product with ID %d", id),
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":     "success",
			"statusCode": 200,
			"message":    fmt.Sprintf("Product with ID %d deleted successfully", id),
		})
	}
}
