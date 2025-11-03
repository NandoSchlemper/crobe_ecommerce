package main

import (
	"context"
	"crobe-ecommerce/app/backend/api/routes"
	"crobe-ecommerce/app/backend/pkg/product"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	godotenv.Load()
	db, client, err := databaseConnection()

	if err != nil {
		log.Fatal("Database Connection Error $s", err)
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
		return ctx.Send([]byte("Welcome to the clean-architecture mongo product shop!"))
	})
	api := app.Group("/api")
	routes.ProductRouter(api, productService)
	log.Fatal(app.Listen(":8080"))
}

func databaseConnection() (*mongo.Database, *mongo.Client, error) {
	uri := os.Getenv("MONGO_DB_URL")

	if uri == "" {
		return nil, nil, errors.New("erro ao carregar a variavel de ambiente do DB")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, nil, err
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, nil, err
	}

	db := client.Database("crobedatabase")
	return db, client, nil
}
