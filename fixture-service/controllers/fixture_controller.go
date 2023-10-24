package controllers

import (
	"context"
	"fixture-service/models"
	"fixture-service/utils"
)

func CreateFixture(fixture *models.Fixture) error {
	_, err := utils.DB.Collection("fixtures").InsertOne(context.TODO(), fixture)
	return err
}
