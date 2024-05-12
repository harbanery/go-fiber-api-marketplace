package routes

import (
	"gofiber-marketplace/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	// Product Routes
	app.Get("/products", controllers.GetAllProduct)
	app.Get("/product/:id", controllers.GetDetailProduct)
	app.Post("/product", controllers.CreateProduct)
	app.Put("/product/:id", controllers.UpdateProduct)
	app.Delete("/product/:id", controllers.DeleteProduct)

	// Category Routes
	app.Get("/categories", controllers.GetAllCategories)
	app.Get("/category/:id", controllers.GetCategoryById)
	app.Post("/category", controllers.CreateCategory)
	app.Put("/category/:id", controllers.UpdateCategory)
	app.Delete("/category/:id", controllers.DeleteCategory)

	// Seller Routes
	app.Post("/sellers/register", controllers.SellerRegister)
	app.Get("/sellers", controllers.GetSellers)
	app.Get("/sellers/:id", controllers.GetDetailSeller)
	app.Put("/sellers/profile/:id", controllers.UpdateSellerProfile)
	app.Delete("/sellers/profile/:id", controllers.DeleteSeller)
}
