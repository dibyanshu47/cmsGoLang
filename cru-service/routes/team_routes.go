package routes

import (
	"cru-service/handlers"

	"github.com/gorilla/mux"
)

func SetTeamRoutes(router *mux.Router) {
	router.HandleFunc("/teams", handlers.CreateTeamHandler).Methods("POST")
	router.HandleFunc("/teams", handlers.GetAllTeamsHandler).Methods("GET")
	router.HandleFunc("/teams/{id}", handlers.GetTeamHandler).Methods("GET")
	router.HandleFunc("/teams/{id}", handlers.UpdateTeamHandler).Methods("PUT")
}
