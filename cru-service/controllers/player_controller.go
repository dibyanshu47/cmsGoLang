package controllers

import (
	"context"
	"cru-service/models"
	"cru-service/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePlayer(player *models.Player) error {
	_, err := utils.DB.Collection("players").InsertOne(context.TODO(), player)
	return err
}

func GetAllPlayers() ([]models.Player, error) {
	var players []models.Player
	cursor, err := utils.DB.Collection("players").Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var player models.Player
		err := cursor.Decode(&player)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}
	return players, nil
}

func GetPlayer(id string) (models.Player, error) {
	var player models.Player
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return player, err
	}
	filter := bson.M{"_id": objectID}
	err = utils.DB.Collection("players").FindOne(context.Background(), filter).Decode(&player)
	return player, err
}

func UpdatePlayer(id string, updatedFields map[string]interface{}) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}

	// Create a `$set` document with the fields and values to update
	update := bson.M{"$set": updatedFields}

	_, err = utils.DB.Collection("players").UpdateOne(context.TODO(), filter, update)
	return err
}
