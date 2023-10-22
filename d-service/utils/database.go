package utils

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDatabase() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://dibyanshujaiswal1607:devrev123@cluster0.g0y0bwz.mongodb.net/?retryWrites=true&w=majority")

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")
	DB = client.Database("CMSGoLang")
}
