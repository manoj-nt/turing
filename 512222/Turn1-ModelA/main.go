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
	ID      string
	Params  map[string]string
	IsValid map[string]bool
}

// Store in-memory query parameter storage
var (
	store      = make(map[string]QueryParameters)
	storeMutex = sync.RWMutex{}
)

// Date validation layout
const dateLayout = "2006-01-02"

// Email validation regular expression
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// Main function
func main() {
	http.HandleFunc("/parameters", parametersHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// HTTP handler for query parameters
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

// Handler for GET request to retrieve parameters
func handleGet(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	storeMutex.RLock()
	defer storeMutex.RUnlock()
	if params, found := store[id]; found {
		json.NewEncoder(w).Encode(params)
	} else {
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

// Handler for POST request to store and validate parameters
func handlePost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	r.ParseForm()
	params := make(map[string]string)
	isValid := make(map[string]bool)
	for key, value := range r.Form {
		params[key] = strings.Join(value, ", ")
		isValid[key] = validateParameter(key, params[key])
	}
	storeMutex.Lock()
	store[id] = QueryParameters{ID: id, Params: params, IsValid: isValid}
	storeMutex.Unlock()
	w.WriteHeader(http.StatusCreated)
}

// Handler for DELETE request to remove parameters
func handleDelete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	storeMutex.Lock()
	delete(store, id)
	storeMutex.Unlock()
	w.WriteHeader(http.StatusNoContent)
}

// Function to validate parameters
func validateParameter(name, value string) bool {
	switch name {
	case "email":
		return emailRegex.MatchString(value)
	case "date":
		_, err := time.Parse(dateLayout, value)
		return err == nil
	default:
		return false
	}
}

// Explain performance considerations
/*
1. Optimal Storage: We are using a sync.RWMutex for safe concurrent access to the in-memory map `store`.
2. Concurrency Handling: By using a read-write lock, we are ensuring that readers do not block each other while locks are shared. Writings will lock the map ensuring that only one write occurs at a time, avoiding race conditions.
3. Compatibility: The logic supports generic query parameters allowing flexibility in use.
4. Performance: HTTP request handlers are inherently concurrent; the mutex ensures the in-memory map's integrity across threads.
*/
