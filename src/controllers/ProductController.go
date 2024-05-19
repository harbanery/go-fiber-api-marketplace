package controllers

import (
	"fmt"
	"gofiber-marketplace/src/helpers"
	"gofiber-marketplace/src/middlewares"
	"gofiber-marketplace/src/models"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllProduct(c *fiber.Ctx) error {
	keyword := c.Query("search")
	sort := helpers.GetSortParams(c.Query("sorting"), c.Query("orderBy"))
	page, limit, offset := helpers.GetPaginationParams(c.Query("limit"), c.Query("page"))
	totalData := models.CountData(keyword)
	totalPage := math.Ceil(float64(totalData) / float64(limit))

	products := models.SelectAllProducts(keyword, sort, limit, offset)
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
			"photo":         product.Image,
			"rating":        product.Rating,
			"price":         product.Price,
			"size":          product.Size,
			"color":         product.Color,
			"stock":         product.Stock,
			"condition":     product.Condition,
			"desc":          product.Description,
		}
	}

	// return c.JSON(products)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "success",
		"statusCode":  200,
		"data":        resultProducts,
		"currentPage": page,
		"limit":       limit,
		"totalData":   totalData,
		"totalPage":   totalPage,
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
		"photo":         product.Image,
		"rating":        product.Rating,
		"price":         product.Price,
		"size":          product.Size,
		"color":         product.Color,
		"stock":         product.Stock,
		"condition":     product.Condition,
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
	auth := middlewares.UserLocals(c)
	if role := auth["role"].(string); role != "seller" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":     "forbidden",
			"statusCode": 403,
			"message":    "Incorrect role",
		})
	}

	var newProduct models.Product

	if err := c.BodyParser(&newProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid request body",
		})
	}

	product := middlewares.XSSMiddleware(&newProduct).(*models.Product)

	if errors := helpers.StructValidation(product); len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":     "unprocessable entity",
			"statusCode": 422,
			"message":    "Validation failed",
			"errors":     errors,
		})
	}

	if category := models.SelectCategoryById(int(product.CategoryID)); category.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Category not found",
		})
	}

	if seller := models.SelectSellerById(int(product.SellerID)); seller.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Seller not found",
		})
	}

	if err := models.CreateProduct(product); err != nil {
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
	auth := middlewares.UserLocals(c)
	if role := auth["role"].(string); role != "seller" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":     "forbidden",
			"statusCode": 403,
			"message":    "Incorrect role",
		})
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid ID format",
		})
	}

	if product := models.SelectProductById(id); product.ID == 0 {
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

	product := middlewares.XSSMiddleware(&updatedProduct).(*models.Product)

	if errors := helpers.StructValidation(product); len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":     "unprocessable entity",
			"statusCode": 422,
			"message":    "Validation failed",
			"errors":     errors,
		})
	}

	if category := models.SelectCategoryById(int(product.CategoryID)); category.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Category not found",
		})
	}

	if seller := models.SelectSellerById(int(product.SellerID)); seller.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Seller not found",
		})
	}

	if err := models.UpdateProduct(id, product); err != nil {
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
	auth := middlewares.UserLocals(c)
	if role := auth["role"].(string); role != "seller" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":     "forbidden",
			"statusCode": 403,
			"message":    "Incorrect role",
		})
	}

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
