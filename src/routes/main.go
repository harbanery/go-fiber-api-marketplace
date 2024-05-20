package routes

import (
	"gofiber-marketplace/src/controllers"
	"gofiber-marketplace/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	// Product Routes
	app.Get("/products", controllers.GetAllProduct)
	app.Get("/product/:id", controllers.GetDetailProduct)
	app.Post("/product", middlewares.JWTMiddleware(), controllers.CreateProduct)
	app.Put("/product/:id", middlewares.JWTMiddleware(), controllers.UpdateProduct)
	app.Delete("/product/:id", middlewares.JWTMiddleware(), controllers.DeleteProduct)

	// Category Routes
	app.Get("/categories", controllers.GetAllCategories)
	app.Get("/category/:id", controllers.GetCategoryById)
	app.Post("/category", middlewares.JWTMiddleware(), controllers.CreateCategory)
	app.Put("/category/:id", middlewares.JWTMiddleware(), controllers.UpdateCategory)
	app.Delete("/category/:id", middlewares.JWTMiddleware(), controllers.DeleteCategory)

	// Seller Routes
	app.Get("/sellers", middlewares.JWTMiddleware(), controllers.GetSellers)
	app.Get("/sellers/:id", middlewares.JWTMiddleware(), controllers.GetDetailSeller)
	app.Get("/seller/profile", middlewares.JWTMiddleware(), controllers.GetSellerProfile)
	app.Put("/seller/profile", middlewares.JWTMiddleware(), controllers.UpdateSellerProfile)
	app.Delete("/seller/profile", middlewares.JWTMiddleware(), controllers.DeleteSeller)
	app.Put("/seller/profile/photo", middlewares.JWTMiddleware(), controllers.UpdateSellerProfilePhoto)

	// Customer Routes
	app.Get("/customers", middlewares.JWTMiddleware(), controllers.GetCustomers)
	app.Get("/customers/:id", middlewares.JWTMiddleware(), controllers.GetDetailCustomer)
	app.Get("/customer/profile", middlewares.JWTMiddleware(), controllers.GetCustomerProfile)
	app.Put("/customer/profile", middlewares.JWTMiddleware(), controllers.UpdateCustomerProfile)
	app.Delete("/customer/profile", middlewares.JWTMiddleware(), controllers.DeleteCustomer)

	// User/Auth Routes
	app.Post("/register", controllers.RegisterUser)
	app.Post("/login", controllers.LoginUser)
	app.Post("/refreshToken", controllers.CreateRefreshToken)

	// Address Routes
	app.Get("/addresses", middlewares.JWTMiddleware(), controllers.GetAddresses)
	app.Post("/address", middlewares.JWTMiddleware(), controllers.CreateAddress)
	app.Put("/address/:id", middlewares.JWTMiddleware(), controllers.UpdateAddress)
	app.Delete("/address/:id", middlewares.JWTMiddleware(), controllers.DeleteAddress)

	// Upload Routes
	app.Post("/upload", controllers.UploadFile)
	app.Post("/uploadServer", controllers.UploadFileServer)
}
