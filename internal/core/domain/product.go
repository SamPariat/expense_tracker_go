package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name          string             `json:"name" bson:"name"`
	Description   *string            `json:"description" bson:"description"`
	Price         float64            `json:"price" bson:"price"`
	Source        string             `json:"source" bson:"source"`
	ImageUrl      *string            `json:"imageUrl" validate:"omitempty" bson:"imageUrl,omitempty"`
	ModeOfPayment string             `json:"modeOfPayment" bson:"modeOfPayment"`
	CreatedAt     primitive.DateTime `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt     primitive.DateTime `json:"updatedAt" bson:"updatedAt,omitempty"`
}
