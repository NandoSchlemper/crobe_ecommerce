package handlers

import (
	"crobe-ecommerce/app/backend/api/presenter"
	"crobe-ecommerce/app/backend/pkg/entities"
	"crobe-ecommerce/app/backend/pkg/product"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func AddProduct(service product.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqBody entities.Product
		err := c.BodyParser(&reqBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ProductErrorResponse(err))
		}

		if reqBody.Name == "" || reqBody.Description == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ProductErrorResponse(errors.New(
				"please specify title and author")))
		}
		result, err := service.InsertProduct(&reqBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ProductErrorResponse(err))
		}
		return c.JSON(presenter.ProductSuccessResponse(result))
	}
}

func GetProduct(service product.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var parameters string
		err := c.ParamsParser(parameters)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ProductErrorResponse(err))
		}

		result, err := service.GetProduct(parameters)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ProductErrorResponse(err))
		}

		return c.JSON(presenter.ProductSuccessResponse(result))
	}
}
