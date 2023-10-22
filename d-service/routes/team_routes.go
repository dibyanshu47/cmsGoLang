package routes

import (
	"d-service/handlers"

	"github.com/gorilla/mux"
)

func SetTeamRoutes(router *mux.Router) {
	router.HandleFunc("/teams/{id}", handlers.DeleteTeamHandler).Methods("DELETE")
}
