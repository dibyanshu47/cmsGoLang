package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Team struct {
	Id          primitive.ObjectID   `json:"_id" bson:"_id,omitempty"`
	Name        string               `json:"name" bson:"name" required:"true"`
	Players     []primitive.ObjectID `json:"players" bson:"players" required:"true"`
	Captain     primitive.ObjectID   `json:"captain" bson:"captain" required:"true"`
	ViceCaptain primitive.ObjectID   `json:"viceCaptain" bson:"viceCaptain" required:"true"`
}
