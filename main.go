package main

import (
	"context"
	"crobe-ecommerce/app/backend/pkg/database"
	"crobe-ecommerce/app/backend/pkg/product"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db, client, err := database.DatabaseConnection()

	if err != nil {
		log.Fatal("Database Connection Error: ", err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	fmt.Println("Database connection success!")

	productCollection := db.Collection("products")
	productRepo := product.NewRepo(productCollection)
	productService := product.NewService(productRepo)

	app := fiber.New()
	// app.Get("/docs/*", scalar.New())
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Crobe E-commerce API is running"))
	})
	api := app.Group("/api")
	product.ProductRouter(api, productService)
	log.Fatal(app.Listen(":8080"))
}
