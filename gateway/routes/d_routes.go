package routes

import (
	"gateway/handlers"

	"github.com/gorilla/mux"
)

func SetDRoutes(router *mux.Router) {
	router.HandleFunc("/d/{entity}/{id}", handlers.DHandler).Methods("DELETE")
}
