package controllers

import (
	"gofiber-marketplace/src/models"
	"reflect"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetAllCategories(c *fiber.Ctx) error {
	categories := models.SelectAllCategories()
	if len(categories) == 0 {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"status":     "no content",
			"statusCode": 202,
			"message":    "Category is empty. You should create product",
		})
	}

	resultCategories := make([]map[string]interface{}, len(categories))
	for i, category := range categories {
		products := make([]map[string]interface{}, len(category.Products))
		for j, product := range category.Products {
			products[j] = map[string]interface{}{
				"id":         product.ID,
				"created_at": product.CreatedAt,
				"updated_at": product.UpdatedAt,
				"name":       product.Name,
				"price":      product.Price,
				"photo":      product.URLImage,
				"size":       product.Size,
				"color":      product.Color,
				"rating":     product.Rating,
			}
		}

		resultCategories[i] = map[string]interface{}{
			"id":         category.ID,
			"created_at": category.CreatedAt,
			"updated_at": category.UpdatedAt,
			"name":       category.Name,
			"image":      category.URLImage,
			"slug":       category.Slug,
			"products":   products,
		}
	}

	// return c.JSON(categories)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"statusCode": 200,
		"data":       resultCategories,
	})
}

func GetCategoryById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid ID format",
		})
	}

	category := models.SelectCategoryById(id)
	if category.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Category not found",
		})
	}

	products := make([]map[string]interface{}, len(category.Products))
	for i, product := range category.Products {
		products[i] = map[string]interface{}{
			"id":         product.ID,
			"created_at": product.CreatedAt,
			"updated_at": product.UpdatedAt,
			"name":       product.Name,
			"price":      product.Price,
			"photo":      product.URLImage,
			"size":       product.Size,
			"color":      product.Color,
			"rating":     product.Rating,
		}
	}

	resultCategory := map[string]interface{}{
		"id":         category.ID,
		"created_at": category.CreatedAt,
		"updated_at": category.UpdatedAt,
		"name":       category.Name,
		"image":      category.URLImage,
		"slug":       category.Slug,
		"products":   products,
	}

	// return c.JSON(category)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"statusCode": 200,
		"data":       resultCategory,
	})
}

func CreateCategory(c *fiber.Ctx) error {
	var category models.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid request body",
		})
	}

	if reflect.TypeOf(category.Name).Kind() != reflect.String {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Category name must be string",
		})
	} else if category.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Category name cannot be empty",
		})
	}

	if category.Slug == "" || category.Slug != strings.ReplaceAll(strings.ToLower(category.Name), " ", "") {
		category.Slug = strings.ReplaceAll(strings.ToLower(category.Name), " ", "")
	}

	existingCategory := models.SelectCategoryBySlug(category.Slug)
	if existingCategory.ID != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Category with this name or slug already exists",
		})
	}

	if err := models.CreateCategory(&category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    "Failed to create category",
		})
	} else {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"status":     "success",
			"statusCode": 200,
			"message":    "Category created successfully",
		})
	}
}

func UpdateCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid ID format",
		})
	}

	category := models.SelectCategoryById(id)
	if category.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Category not found",
		})
	}

	var updatedCategory models.Category

	if err := c.BodyParser(&updatedCategory); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid request body",
		})

	}

	if reflect.TypeOf(updatedCategory.Name).Kind() != reflect.String {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Category name must be string",
		})
	} else if updatedCategory.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Category name cannot be empty",
		})
	}

	// if updatedCategory.URLImage == "" {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"status":     "bad request",
	// 		"statusCode": 400,
	// 		"message":    "Category photo cannot be empty",
	// 	})
	// }

	// slug harus sama dengan nama dengan huruf kecil dan tanpa spasi
	if updatedCategory.Slug == "" || updatedCategory.Slug != strings.ReplaceAll(strings.ToLower(updatedCategory.Name), " ", "") {
		updatedCategory.Slug = strings.ReplaceAll(strings.ToLower(updatedCategory.Name), " ", "")
	}

	existingCategory := models.SelectCategoryBySlug(updatedCategory.Slug)
	if existingCategory.ID != 0 && existingCategory.ID != updatedCategory.ID {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Category with this name or slug already exists",
		})
	}

	if err := models.UpdateCategory(id, &updatedCategory); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    "Failed to update category",
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":     "success",
			"statusCode": 200,
			"message":    "Category updated successfully",
		})
	}
}

func DeleteCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid ID format",
		})
	}

	category := models.SelectCategoryById(id)
	if category.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Category not found",
		})
	}

	if err := models.DeleteCategory(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    "Failed to delete category",
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":     "success",
			"statusCode": 200,
			"message":    "Category deleted successfully",
		})
	}
}
