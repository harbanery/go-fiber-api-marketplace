package controllers

import (
	"gofiber-marketplace/src/helpers"
	"gofiber-marketplace/src/middlewares"
	"gofiber-marketplace/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetSellers(c *fiber.Ctx) error {
	sellers := models.SelectAllSellers()
	if len(sellers) == 0 {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"status":     "no content",
			"statusCode": 202,
			"message":    "Seller is empty",
		})
	}

	resultSellers := make([]map[string]interface{}, len(sellers))
	for i, seller := range sellers {
		products := make([]map[string]interface{}, len(seller.Products))
		for j, product := range seller.Products {
			var categoryName string
			if product.CategoryID != 0 {
				category := models.SelectCategoryById(int(product.CategoryID))
				if category.ID != 0 {
					categoryName = category.Name
				}
			}

			products[j] = map[string]interface{}{
				"id":            product.ID,
				"created_at":    product.CreatedAt,
				"updated_at":    product.UpdatedAt,
				"name":          product.Name,
				"price":         product.Price,
				"size":          product.Size,
				"color":         product.Color,
				"photo":         product.Image,
				"rating":        product.Rating,
				"category_name": categoryName,
			}
		}

		resultSellers[i] = map[string]interface{}{
			"id":         seller.ID,
			"created_at": seller.CreatedAt,
			"updated_at": seller.UpdatedAt,
			"name":       seller.Name,
			"user_id":    seller.User.ID,
			"email":      seller.User.Email,
			"photo":      seller.Image,
			"phone":      seller.Phone,
			"desc":       seller.Description,
			"role":       seller.User.Role,
			"products":   products,
		}
	}

	// return c.JSON(categories)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"statusCode": 200,
		"data":       resultSellers,
	})
}

func GetDetailSeller(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid ID format",
		})
	}

	seller := models.SelectSellerById(id)
	if seller.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Seller not found",
		})
	}

	if seller.User.Role != "seller" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Role of this user is customer or not seller",
		})
	}

	products := make([]map[string]interface{}, len(seller.Products))
	for i, product := range seller.Products {
		var categoryName string
		if product.CategoryID != 0 {
			category := models.SelectCategoryById(int(product.CategoryID))
			if category.ID != 0 {
				categoryName = category.Name
			}
		}

		products[i] = map[string]interface{}{
			"id":            product.ID,
			"created_at":    product.CreatedAt,
			"updated_at":    product.UpdatedAt,
			"name":          product.Name,
			"price":         product.Price,
			"size":          product.Size,
			"color":         product.Color,
			"photo":         product.Image,
			"rating":        product.Rating,
			"category_name": categoryName,
		}
	}

	resultSeller := map[string]interface{}{
		"id":         seller.ID,
		"created_at": seller.CreatedAt,
		"updated_at": seller.UpdatedAt,
		"name":       seller.Name,
		"user_id":    seller.User.ID,
		"email":      seller.User.Email,
		"photo":      seller.Image,
		"phone":      seller.Phone,
		"desc":       seller.Description,
		"role":       seller.User.Role,
		"products":   products,
	}

	// return c.JSON(product)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"statusCode": 200,
		"data":       resultSeller,
	})
}

func GetSellerProfile(c *fiber.Ctx) error {
	auth := middlewares.UserLocals(c)
	if role := auth["role"].(string); role != "seller" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":     "forbidden",
			"statusCode": 403,
			"message":    "Incorrect role",
		})
	}

	id, ok := auth["id"].(float64)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid ID format",
		})
	}

	seller := models.SelectSellerByUserId(int(id))
	if seller.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Seller not found",
		})
	}

	if seller.User.Role != "seller" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Role of this user is customer or not seller",
		})
	}

	products := make([]map[string]interface{}, len(seller.Products))
	for i, product := range seller.Products {
		var categoryName string
		if product.CategoryID != 0 {
			category := models.SelectCategoryById(int(product.CategoryID))
			if category.ID != 0 {
				categoryName = category.Name
			}
		}

		products[i] = map[string]interface{}{
			"id":            product.ID,
			"created_at":    product.CreatedAt,
			"updated_at":    product.UpdatedAt,
			"name":          product.Name,
			"price":         product.Price,
			"size":          product.Size,
			"color":         product.Color,
			"photo":         product.Image,
			"rating":        product.Rating,
			"category_name": categoryName,
		}
	}

	resultSeller := map[string]interface{}{
		"id":         seller.ID,
		"created_at": seller.CreatedAt,
		"updated_at": seller.UpdatedAt,
		"name":       seller.Name,
		"user_id":    seller.User.ID,
		"email":      seller.User.Email,
		"photo":      seller.Image,
		"phone":      seller.Phone,
		"desc":       seller.Description,
		"role":       seller.User.Role,
		"products":   products,
	}

	// return c.JSON(product)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"statusCode": 200,
		"data":       resultSeller,
	})
}

type SellerProfile struct {
	Name        string `json:"name" validate:"required,max=50"`
	Email       string `json:"email" validate:"required,email"`
	Image       string `json:"image" validate:"required"`
	Phone       string `json:"phone" validate:"required,numeric,max=15"`
	Description string `json:"description" validate:"required"`
}

func UpdateSellerProfile(c *fiber.Ctx) error {
	var profileData SellerProfile

	auth := middlewares.UserLocals(c)
	if role := auth["role"].(string); role != "seller" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":     "forbidden",
			"statusCode": 403,
			"message":    "Incorrect role",
		})
	}

	id, ok := auth["id"].(float64)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid ID format",
		})
	}

	seller := models.SelectSellerByUserId(int(id))
	if seller.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Seller not found",
		})
	}

	if seller.User.Role != "seller" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Role of this user is not seller",
		})
	}

	if seller.UserID == 0 || seller.User.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Seller not found the user",
		})
	}

	if err := c.BodyParser(&profileData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid request body",
		})
	}

	user := middlewares.XSSMiddleware(&profileData).(*SellerProfile)
	if errors := helpers.StructValidation(user); len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":     "unprocessable entity",
			"statusCode": 422,
			"message":    "Validation failed",
			"errors":     errors,
		})
	}

	if existUser := models.SelectUserbyEmail(user.Email); existUser.ID != 0 && existUser.Email != user.Email {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Email already exists",
		})
	}

	updatedUser := models.User{
		Email: user.Email,
	}

	updatedSeller := models.Seller{
		Name:        user.Name,
		Image:       user.Image,
		Phone:       user.Phone,
		Description: user.Description,
	}

	if err := models.UpdateUser(int(seller.User.ID), &updatedUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    "Failed to update user",
		})
	}

	if err := models.UpdateSeller(int(id), &updatedSeller); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    "Failed to update seller",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"statusCode": 200,
		"message":    "Profile updated successfully",
	})

}

func DeleteSeller(c *fiber.Ctx) error {
	auth := middlewares.UserLocals(c)
	if role := auth["role"].(string); role != "seller" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":     "forbidden",
			"statusCode": 403,
			"message":    "Incorrect role",
		})
	}

	id, ok := auth["id"].(float64)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid ID format",
		})
	}

	seller := models.SelectSellerByUserId(int(id))
	if seller.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Seller not found",
		})
	}

	if seller.User.Role != "seller" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Role of this user is customer or not seller",
		})
	}

	if seller.UserID == 0 || seller.User.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Seller not found the user",
		})
	}

	if err := models.DeleteSeller(int(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    "Failed to delete seller",
		})
	}

	if err := models.DeleteUser(int(seller.User.ID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    "Failed to delete user",
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":     "success",
			"statusCode": 200,
			"message":    "User deleted successfully",
		})
	}
}
