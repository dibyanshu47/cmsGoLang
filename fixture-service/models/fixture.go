package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Fixture struct {
	Id      primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Matches []Match            `json:"matches" bson:"matches" required:"true"`
}

type ResultChoice string

const (
	Win       ResultChoice = "Win"
	Lose      ResultChoice = "Lose"
	NotPlayed ResultChoice = "NotPlayed"
)

type Match struct {
	Team1  primitive.ObjectID `json:"team1" bson:"team1"`
	Team2  primitive.ObjectID `json:"team2" bson:"team2"`
	Result ResultChoice       `json:"winner" bson:"winner"`
}
