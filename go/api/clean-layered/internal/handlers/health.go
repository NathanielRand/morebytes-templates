package handlers

import (
	"encoding/json"
	"net/http"
)

// HealthHandler is a handler for the /health endpoint.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	// Create and populate the response object
	response := map[string]string{"status": "OK"}

	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode the response object as JSON and write it to the response
	// and return an error if the encoding fails.
	if err := json.NewEncoder(w).Encode(response); err != nil {
		// Set the response content type 
		w.Header().Set("Content-Type", "text/plain")
		// Set the response status code to 500 Internal Server Error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}