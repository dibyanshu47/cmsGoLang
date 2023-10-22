package main

import (
	"d-service/routes"
	"d-service/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	utils.ConnectDatabase()

	router := mux.NewRouter()

	routes.SetPlayerRoutes(router)
	routes.SetTeamRoutes(router)

	http.Handle("/", router)

	fmt.Println("Server is running on port 8082")
	http.ListenAndServe(":8082", nil)
}
