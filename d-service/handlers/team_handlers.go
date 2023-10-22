package handlers

import (
	"d-service/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteTeamHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamId := vars["id"]

	if err := controllers.DeleteTeam(teamId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
