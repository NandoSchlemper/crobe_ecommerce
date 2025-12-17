package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository interface {
	CreateProduct(product *Product) (*Product, error)
	FetchProduct(id string) (*Product, error)
	DeleteProduct(product *DeleteRequest) error
}

type repository struct {
	Collection *mongo.Collection
}

func (r *repository) FetchProduct(id string) (*Product, error) {
	var product Product
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
func (r *repository) CreateProduct(product *Product) (*Product, error) {
	product.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *repository) DeleteProduct(product *DeleteRequest) error {
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
