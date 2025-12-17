package product

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductPresenter struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Description string             `json:"description"`
}

// Singular SuccessResponse that will be passed in the response by Handler
func ProductSuccessResponse(data *Product) *fiber.Map {
	product := Product{
		ID:          data.ID,
		Description: data.Description,
	}
	return &fiber.Map{
		"status": true,
		"data":   product,
		"error":  nil,
	}
}

// List of SuccessResponse that will be passed in the response by Handler
func ProductsSuccessResponse(data *[]ProductPresenter) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

// Basically the error response by Handler
func ProductErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
