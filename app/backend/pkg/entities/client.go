package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Client struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email    string             `json:"email" bson:"email"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Address  interface{}        `json:"address"` // preciso pesquisar em qual formato guardar os addresses
}
