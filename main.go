// Import required packages
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// In-memory store for the key-value pairs
var store = make(map[string]string)

// Handler function for /set route
// This function is called whenever a POST request is made to /set
// It expects a JSON object in the request body
// The keys and values from the JSON object are stored in the in-memory store
func setHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON object from the request body
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		// If there is an error, return a HTTP status 400 Bad Request
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Iterate over the keys and values from the JSON object and store them
	for key, value := range data {
		store[key] = value
	}

	// If successful, return a HTTP status 204 No Content
	w.WriteHeader(http.StatusNoContent)
}

// Handler function for /get/{key} route
// This function is called whenever a GET request is made to /get/{key}
// It returns the value for the key as a JSON object
func getHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the key from the URL path parameters
	key := mux.Vars(r)["key"]
	// Lookup the key in the store
	value, ok := store[key]
	if !ok {
		// If the key is not found, return a HTTP status 404 Not Found
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}

	// If the key is found, return the value as a JSON object
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{key: value})
}

func main() {
	// Create a new Gorilla Mux router
	r := mux.NewRouter()

	// Register the handler functions for the routes
	r.HandleFunc("/set", setHandler).Methods("POST")
	r.HandleFunc("/get/{key}", getHandler).Methods("GET")
	// Print REDIS server start message to the terminal
	fmt.Println("Starting REDIS server on port 8001")

	// Start the HTTP server and listen on port 8000
	// If there is an error, log the error and exit
	log.Fatal(http.ListenAndServe(":8001", r))
}
