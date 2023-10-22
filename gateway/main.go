package main

import (
	"gateway/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// routes.SetupRoutes(router)
	routes.SetCruRoutes(router)
	routes.SetDRoutes(router)

	http.Handle("/", router)

	log.Println("Starting API Gateway on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
