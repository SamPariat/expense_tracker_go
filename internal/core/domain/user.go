package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	Income    float64            `json:"income" bson:"income"`
	ImageUrl  *string            `json:"imageUrl" validate:"omitempty" bson:"imageUrl,omitempty"`
	CreatedAt primitive.DateTime `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt primitive.DateTime `json:"updatedAt" bson:"updatedAt,omitempty"`
}
