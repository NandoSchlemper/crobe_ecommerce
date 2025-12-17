package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Client struct {
	ID       primitive.ObjectID `json:"id"`
	Email    string             `json:"email"`
	Username string             `json:"username"`
	Password string             `json:"password"`
	Address  interface{}        `json:"address"`
}
