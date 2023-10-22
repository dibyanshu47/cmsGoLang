package controllers

import (
	"context"
	"cru-service/models"
	"cru-service/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTeam(team *models.Team) error {
	_, err := utils.DB.Collection("teams").InsertOne(context.TODO(), team)
	return err
}

func GetAllTeams() ([]models.Team, error) {
	var teams []models.Team
	cursor, err := utils.DB.Collection("teams").Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &teams); err != nil {
		return nil, err
	}
	return teams, nil
}

func GetTeam(id string) (models.Team, error) {
	var team models.Team
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return team, err
	}
	filter := bson.M{"_id": objectID}
	err = utils.DB.Collection("teams").FindOne(context.Background(), filter).Decode(&team)
	return team, err
}

func UpdateTeam(id string, updatedFields map[string]interface{}) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}

	// Create a `$set` document with the fields and values to update
	update := bson.M{"$set": updatedFields}

	_, err = utils.DB.Collection("teams").UpdateOne(context.TODO(), filter, update)
	return err
}
