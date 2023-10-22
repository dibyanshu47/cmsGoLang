package routes

import (
	"gateway/handlers"

	"github.com/gorilla/mux"
)

func SetCruRoutes(router *mux.Router) {
	router.HandleFunc("/cru/{entity}", handlers.CruHandler).Methods("GET")
	router.HandleFunc("/cru/{entity}", handlers.CruHandler).Methods("POST")
	router.HandleFunc("/cru/{entity}/{id}", handlers.CruHandler).Methods("PUT")
}
