package controllers

import (
	"context"
	"d-service/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeletePlayer(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}

	_, err = utils.DB.Collection("players").DeleteOne(context.TODO(), filter)
	if err != nil && err == mongo.ErrNoDocuments {
		return nil // Player not found, no error
	}

	return err
}
