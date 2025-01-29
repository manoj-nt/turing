package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

// Struct to store query parameters and corresponding validation results
type QueryParameters struct {
	ID      string            `json:"id"`
	Params  map[string]string `json:"params"`
	IsValid map[string]bool   `json:"is_valid"`
}

// In-memory storage of query parameters
var (
	store      = make(map[string]QueryParameters)
	storeMutex = sync.RWMutex{}
)

// Date validation layout
const dateLayout = "2006-01-02"

// Email validation regular expression
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// Main function to set up the HTTP server
func main() {
	http.HandleFunc("/parameters", parametersHandler)
	log.Println("Server started on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// HTTP handler to process query parameters
func parametersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGet(w, r)
	case http.MethodPost:
		handlePost(w, r)
	case http.MethodDelete:
		handleDelete(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Handler for GET requests to retrieve parameters by ID
func handleGet(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	storeMutex.RLock()
	defer storeMutex.RUnlock()

	if params, found := store[id]; found {
		respondWithJSON(w, http.StatusOK, params)
	} else {
		http.Error(w, "Parameters not found", http.StatusNotFound)
	}
}

// Handler for POST requests to store and validate parameters
func handlePost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	params := make(map[string]string)
	isValid := make(map[string]bool)

	// Validate parameters and store them
	for key, values := range r.Form {
		paramValue := strings.Join(values, ", ")
		params[key] = paramValue
		isValid[key] = validateParameter(key, paramValue)
	}

	// Store the validated parameters in the in-memory map
	storeMutex.Lock()
	store[id] = QueryParameters{ID: id, Params: params, IsValid: isValid}
	storeMutex.Unlock()

	log.Printf("Stored parameters for ID %s", id)
	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Parameters stored successfully"})
}

// Handler for DELETE requests to remove stored parameters by ID
func handleDelete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	storeMutex.Lock()
	defer storeMutex.Unlock()

	if _, found := store[id]; found {
		delete(store, id)
		log.Printf("Deleted parameters for ID %s", id)
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "Parameters not found", http.StatusNotFound)
	}
}

// Function to validate query parameters based on their name
func validateParameter(name, value string) bool {
	switch name {
	case "email":
		return emailRegex.MatchString(value)
	case "date":
		_, err := time.Parse(dateLayout, value)
		return err == nil
	default:
		// Add other validation rules for different parameters if needed
		return true
	}
}

// Helper function to send JSON responses
func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
