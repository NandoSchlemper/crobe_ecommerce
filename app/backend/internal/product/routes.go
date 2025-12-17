package product

import (
	"github.com/gofiber/fiber/v2"
)

func ProductRouter(app fiber.Router, service Service) {
	app.Post("/product", AddProduct(service))
	app.Get("/product/:id", GetProduct(service)) // individual Products
	app.Get("/product")                          // multiple products
}
