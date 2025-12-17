package database

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DatabaseConnection() (*mongo.Database, *mongo.Client, error) {
	uri := os.Getenv("MONGO_DB_URL")
	name := os.Getenv("MONGO_DB_NAME")

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

	db := client.Database(name)
	return db, client, nil
}
