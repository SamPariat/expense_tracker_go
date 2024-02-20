package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Travel struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Amount      float64            `json:"amount" bson:"amount"`
	PaidTo      string             `json:"paidTo" bson:"paidTo"`
	Source      string             `json:"source" bson:"source"`
	Destination string             `json:"destination" bson:"destination"`
	StartTime   primitive.DateTime `json:"startTime" bson:"startTime"`
	EndTime     primitive.DateTime `json:"endTime" bson:"endTime"`
	CreatedAt   primitive.DateTime `json:"createdAt" bson:"createdAtomitempty"`
	UpdatedAt   primitive.DateTime `json:"updatedAt" bson:"updatedAt,omitempty"`
}
