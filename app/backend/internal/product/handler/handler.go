package handler

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func AddProduct(service Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqBody Product
		err := c.BodyParser(&reqBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(ProductErrorResponse(err))
		}

		if reqBody.Name == "" || reqBody.Description == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(ProductErrorResponse(errors.New(
				"please specify title and author")))
		}
		result, err := service.InsertProduct(&reqBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(ProductErrorResponse(err))
		}
		return c.JSON(ProductSuccessResponse(result))
	}
}

func GetProduct(service Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var parameters string
		err := c.ParamsParser(parameters)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(ProductErrorResponse(err))
		}

		result, err := service.GetProduct(parameters)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(ProductErrorResponse(err))
		}

		return c.JSON(ProductSuccessResponse(result))
	}
}
