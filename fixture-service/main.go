package main

import (
	"fixture-service/routes"
	"fixture-service/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	utils.ConnectDatabase()

	router := mux.NewRouter()

	routes.SetFixtureRoutes(router)

	http.Handle("/", router)

	fmt.Println("Server is running on port 8083")
	http.ListenAndServe(":8083", nil)
}
