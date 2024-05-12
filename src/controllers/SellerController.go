package controllers

import (
	"gofiber-marketplace/src/models"
	"reflect"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SellerRegister(c *fiber.Ctx) error {
	var newUser models.User
	var newSeller models.Seller
	var registrationData struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&registrationData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid request body",
		})
	}

	if reflect.TypeOf(registrationData.Name).Kind() != reflect.String {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Name must be string",
		})
	} else if registrationData.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Name cannot be empty",
		})
	} else {
		newSeller.Name = registrationData.Name
	}

	if reflect.TypeOf(registrationData.Email).Kind() != reflect.String {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Email must be string",
		})
	} else if registrationData.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Email cannot be empty",
		})
	} else {
		existingUser := models.SelectUserbyEmail(registrationData.Email)
		if existingUser.ID != 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":     "bad request",
				"statusCode": 400,
				"message":    "Email already exists",
			})
		} else {
			newUser.Email = registrationData.Email
		}
	}

	if registrationData.Phone == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Phone cannot be empty",
		})
	} else {
		newSeller.Phone = registrationData.Phone
	}

	if registrationData.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Password cannot be empty",
		})
	} else {
		newUser.Password = registrationData.Password
	}

	if newUser.Role != "seller" {
		newUser.Role = "seller"
	}

	userId, err := models.CreateUser(&newUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    "Failed to create user",
		})
	} else {
		newSeller.UserID = userId
	}

	if err := models.CreateSeller(&newSeller); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    "Failed to create seller",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":     "success",
		"statusCode": 200,
		"message":    "User created successfully",
	})
}

func GetSellers(c *fiber.Ctx) error {
	sellers := models.SelectAllSellers()
	if len(sellers) == 0 {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"status":     "no content",
			"statusCode": 202,
			"message":    "Product is empty. You should create product",
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
				"photo":         product.URLImage,
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
			"photo":      seller.URLImage,
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
	// masih id karena belum ada token/auth
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
			"photo":         product.URLImage,
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
		"photo":      seller.URLImage,
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

func UpdateSellerProfile(c *fiber.Ctx) error {
	var updatedUser models.User
	var updatedSeller models.Seller
	var profileData struct {
		Name        string `json:"name"`
		Email       string `json:"email"`
		Phone       string `json:"phone"`
		Description string `json:"description"`
		URLImage    string `json:"url_image"`
	}

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

	// validasi email
	if reflect.TypeOf(profileData.Email).Kind() != reflect.String {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Email must be string",
		})
	} else if profileData.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Email cannot be empty",
		})
	} else if profileData.Email != seller.User.Email {
		existingUser := models.SelectUserbyEmail(profileData.Email)
		if existingUser.ID != 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":     "bad request",
				"statusCode": 400,
				"message":    "Email already exists",
			})
		} else {
			updatedUser.Email = profileData.Email
		}

		if err := models.UpdateUser(int(seller.User.ID), &updatedUser); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":     "server error",
				"statusCode": 500,
				"message":    "Failed to update user",
			})
		}
	}

	if reflect.TypeOf(profileData.Name).Kind() != reflect.String {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Name must be string",
		})
	} else if profileData.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Name cannot be empty",
		})
	} else {
		updatedSeller.Name = profileData.Name
	}

	if profileData.Phone == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Phone cannot be empty",
		})
	} else {
		updatedSeller.Phone = profileData.Phone
	}

	if profileData.Description == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Description cannot be empty",
		})
	} else {
		updatedSeller.Description = profileData.Description
	}

	// if profileData.URLImage == "" {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"status":     "bad request",
	// 		"statusCode": 400,
	// 		"message":    "URLImage cannot be empty",
	// 	})
	// } else {
	// 	updatedSeller.URLImage = profileData.URLImage
	// }

	if err := models.UpdateSeller(id, &updatedSeller); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    "Failed to update seller",
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":     "success",
			"statusCode": 200,
			"message":    "Seller updated successfully",
		})
	}
}

func DeleteSeller(c *fiber.Ctx) error {
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

	if seller.UserID == 0 || seller.User.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Seller not found the user",
		})
	}

	if err := models.DeleteSeller(id); err != nil {
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
