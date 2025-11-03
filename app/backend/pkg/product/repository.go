package product

import (
	"context"
	"crobe-ecommerce/app/backend/pkg/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository interface {
	CreateProduct(product *entities.Product) (*entities.Product, error)
}

type repository struct {
	Collection *mongo.Collection
}

// CreateProduct implements Repository.
func (r *repository) CreateProduct(book *entities.Product) (*entities.Product, error) {
	book.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func NewRepo(coll *mongo.Collection) Repository {
	return &repository{
		Collection: coll,
	}
}
