package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Owe struct {
	Id           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Amount       float64            `json:"amount" bson:"amount"`
	PersonOwedTo string             `json:"personOwedTo" bson:"personOwedTo"`
	Paid         bool               `json:"paid" bson:"paid"`
	CreatedAt    primitive.DateTime `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt    primitive.DateTime `json:"updatedAt" bson:"updatedAt,omitempty"`
}
