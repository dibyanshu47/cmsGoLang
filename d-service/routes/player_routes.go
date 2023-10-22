package routes

import (
	"d-service/handlers"

	"github.com/gorilla/mux"
)

func SetPlayerRoutes(router *mux.Router) {
	router.HandleFunc("/players/{id}", handlers.DeletePlayerHandler).Methods("DELETE")
}
