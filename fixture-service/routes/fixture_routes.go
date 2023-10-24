package routes

import (
	"fixture-service/handlers"

	"github.com/gorilla/mux"
)

func SetFixtureRoutes(router *mux.Router) {
	router.HandleFunc("/fixtures", handlers.CreateFixtureHandler).Methods("POST")
	// router.HandleFunc("/fixtures", GetFixturesHandler).Methods("GET")
	// router.HandleFunc("/fixtures/{id}", GetFixtureHandler).Methods("GET")
	// router.HandleFunc("/fixtures/{id}", DeleteFixtureHandler).Methods("DELETE")
	// router.HandleFunc("/fixtures/{id}", UpdateFixtureHandler).Methods("PUT")
}
