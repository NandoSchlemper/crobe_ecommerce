package product

import (
	"context"
	"crobe-ecommerce/app/backend/pkg/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository interface {
	CreateProduct(product *entities.Product) (*entities.Product, error)
	FetchProduct(id string) (*entities.Product, error)
	DeleteProduct(product *entities.DeleteRequest) error
}

type repository struct {
	Collection *mongo.Collection
}

func (r *repository) FetchProduct(id string) (*entities.Product, error) {
	var product entities.Product
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	err = r.Collection.FindOne(context.Background(), filter).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// CreateProduct implements Repository.
func (r *repository) CreateProduct(product *entities.Product) (*entities.Product, error) {
	product.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *repository) DeleteProduct(product *entities.DeleteRequest) error {
	objId, err := primitive.ObjectIDFromHex(product.ID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objId}
	_, err = r.Collection.DeleteOne(context.Background(), filter)
	return err
}

func NewRepo(coll *mongo.Collection) Repository {
	return &repository{
		Collection: coll,
	}
}
