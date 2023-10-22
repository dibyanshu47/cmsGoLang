package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func DHandler(w http.ResponseWriter, r *http.Request) {
	entity := mux.Vars(r)["entity"]
	url := "http://localhost:8082/" + entity
	var method string

	switch r.Method {
	case http.MethodDelete:
		// Extract the ID parameter from the URL
		vars := mux.Vars(r)
		id := vars["id"]

		if id != "" {
			url += "/" + id
			method = http.MethodDelete
		} else {
			http.Error(w, "Missing 'id' parameter in the URL", http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	resp, err := makeHTTPRequest(url, method, nil)
	if err != nil {
		http.Error(w, "Error connecting to Hello Service", http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}
