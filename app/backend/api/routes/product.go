package routes

import (
	"crobe-ecommerce/app/backend/api/handlers"
	"crobe-ecommerce/app/backend/pkg/product"

	"github.com/gofiber/fiber/v2"
)

func ProductRouter(app fiber.Router, service product.Service) {
	app.Post("/product", handlers.AddProduct(service))
	app.Get("/product/:id", handlers.GetProduct(service)) // individual Products
	app.Get("/product")                                   // multiple products
}
