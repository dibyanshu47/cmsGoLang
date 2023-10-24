package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Player struct {
	Id            primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name          string             `json:"name" bson:"name" required:"true"`
	Age           int                `json:"age" bson:"age" required:"true"`
	Role          string             `json:"role" bson:"role" required:"true"`
	BattingAvg    float64            `json:"battingAvg" bson:"battingAvg"`
	StrikeRate    float64            `json:"strikeRate" bson:"strikeRate"`
	Economy       float64            `json:"economy" bson:"economy"`
	MatchesPlayed int                `json:"matchesPlayed" bson:"matchesPlayed"`
	TotalRuns     int                `json:"totalRuns" bson:"totalRuns"`
	TotalWickets  int                `json:"totalWickets" bson:"totalWickets"`
}
