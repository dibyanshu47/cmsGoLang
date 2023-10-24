package handlers

import (
	"encoding/json"
	"fixture-service/controllers"
	"fixture-service/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 1 2 3 4 5 6

// 1 2
// 1 3
// 1 4
// 1 5
// 1 6
// 2 3
// 2 4
// 2 5
// 2 6
// 3 4
// 3 5
// 3 6
// 4 5
// 4 6
// 5 6

func generateFixture(teams []primitive.ObjectID) models.Fixture {
	var matches []models.Match

	for i := 0; i < len(teams); i++ {
		for j := i + 1; j < len(teams); j++ {
			matches = append(matches, models.Match{
				Team1:  teams[i],
				Team2:  teams[j],
				Result: models.NotPlayed,
			})
		}
	}

	return models.Fixture{
		Id:      primitive.NewObjectID(),
		Matches: matches,
	}
}

func CreateFixtureHandler(w http.ResponseWriter, r *http.Request) {
	var teams []primitive.ObjectID
	if err := json.NewDecoder(r.Body).Decode(&teams); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newFixture := generateFixture(teams)

	if err := controllers.CreateFixture(&newFixture); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newFixture)
}
