package routes

import (
	"cru-service/handlers"

	"github.com/gorilla/mux"
)

func SetPlayerRoutes(router *mux.Router) {
	router.HandleFunc("/players", handlers.CreatePlayerHandler).Methods("POST")
	router.HandleFunc("/players", handlers.GetAllPlayersHandler).Methods("GET")
	router.HandleFunc("/players/{id}", handlers.GetPlayerHandler).Methods("GET")
	router.HandleFunc("/players/{id}", handlers.UpdatePlayerHandler).Methods("PUT")
}
