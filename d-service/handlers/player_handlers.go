package handlers

import (
	"d-service/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func DeletePlayerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playerID := vars["id"]

	if err := controllers.DeletePlayer(playerID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
