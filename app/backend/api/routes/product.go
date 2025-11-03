package routes

import (
	"crobe-ecommerce/app/backend/api/handlers"
	"crobe-ecommerce/app/backend/pkg/product"

	"github.com/gofiber/fiber/v2"
)

func BookRouter(app fiber.Router, service product.Service) {
	app.Get("/books", handlers.AddProduct(service))
}
