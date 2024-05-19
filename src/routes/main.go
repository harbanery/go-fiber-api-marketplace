package routes

import (
	"gofiber-marketplace/src/controllers"
	"gofiber-marketplace/src/helpers"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	// Product Routes
	app.Get("/products", controllers.GetAllProduct)
	app.Get("/product/:id", controllers.GetDetailProduct)
	app.Post("/product", helpers.JWTMiddleware(), controllers.CreateProduct)
	app.Put("/product/:id", helpers.JWTMiddleware(), controllers.UpdateProduct)
	app.Delete("/product/:id", helpers.JWTMiddleware(), controllers.DeleteProduct)

	// Category Routes
	app.Get("/categories", controllers.GetAllCategories)
	app.Get("/category/:id", controllers.GetCategoryById)
	app.Post("/category", helpers.JWTMiddleware(), controllers.CreateCategory)
	app.Put("/category/:id", helpers.JWTMiddleware(), controllers.UpdateCategory)
	app.Delete("/category/:id", helpers.JWTMiddleware(), controllers.DeleteCategory)

	// Seller Routes
	app.Get("/sellers", helpers.JWTMiddleware(), controllers.GetSellers)
	app.Get("/sellers/:id", helpers.JWTMiddleware(), controllers.GetDetailSeller)
	app.Get("/seller/profile", helpers.JWTMiddleware(), controllers.GetSellerProfile)
	app.Put("/seller/profile", helpers.JWTMiddleware(), controllers.UpdateSellerProfile)
	app.Delete("/seller/profile", helpers.JWTMiddleware(), controllers.DeleteSeller)

	// Customer Routes
	app.Get("/customers", helpers.JWTMiddleware(), controllers.GetCustomers)
	app.Get("/customer/:id", helpers.JWTMiddleware(), controllers.GetDetailCustomer)
	app.Put("/customer/profile/:id", helpers.JWTMiddleware(), controllers.UpdateCustomerProfile)
	app.Delete("/customer/profile/:id", helpers.JWTMiddleware(), controllers.DeleteCustomer)

	// User/Auth Routes
	app.Post("/register", controllers.RegisterUser)
	app.Post("/login", controllers.LoginUser)
	app.Post("/refreshToken", controllers.CreateRefreshToken)

	// Address Routes
	app.Get("/addresses", helpers.JWTMiddleware(), controllers.GetAddresses)
	app.Post("/address", helpers.JWTMiddleware(), controllers.CreateAddress)
	app.Put("/address/:id", helpers.JWTMiddleware(), controllers.UpdateAddress)
	app.Delete("/address/:id", helpers.JWTMiddleware(), controllers.DeleteAddress)

	// Upload Routes
	app.Post("/upload", controllers.UploadFile)
	app.Post("/uploadServer", controllers.UploadFileServer)
}
