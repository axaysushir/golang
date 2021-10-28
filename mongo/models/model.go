package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// func model() {
// 	fmt.Println("Models in go")
// }

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}
