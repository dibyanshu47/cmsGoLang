package handlers

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func makeHTTPRequest(url string, method string, body []byte) ([]byte, error) {
	// Create an HTTP client
	client := &http.Client{}

	// Create a new HTTP request
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	// Set request headers if needed
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBytes, nil
}

func CruHandler(w http.ResponseWriter, r *http.Request) {
	entity := mux.Vars(r)["entity"]
	url := "http://localhost:8081/" + entity
	var method string
	var requestBody []byte // Declare requestBody outside the switch statement

	switch r.Method {
	case http.MethodGet:
		method = http.MethodGet
	case http.MethodPost:
		method = http.MethodPost

		// Capture the request body in a buffer
		requestBodyBuffer := new(bytes.Buffer)
		_, err := requestBodyBuffer.ReadFrom(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Use the captured request body
		requestBody = requestBodyBuffer.Bytes() // Assign to the outer requestBody variable
	case http.MethodPut:
		// Extract the ID parameter from the URL
		vars := mux.Vars(r)
		id := vars["id"]

		if id != "" {
			url += "/" + id
			method = http.MethodPut

			// Capture the request body in a buffer
			requestBodyBuffer := new(bytes.Buffer)
			_, err := requestBodyBuffer.ReadFrom(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Use the captured request body
			requestBody = requestBodyBuffer.Bytes()
		} else {
			http.Error(w, "Missing 'id' parameter in the URL", http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	resp, err := makeHTTPRequest(url, method, requestBody)
	if err != nil {
		http.Error(w, "Error connecting to Cru Service", http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}
