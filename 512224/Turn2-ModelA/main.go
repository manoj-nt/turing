package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the control command from URL query parameters
	queryParams := r.URL.Query()
	action := queryParams.Get("action")    // "play", "pause", "forward", "rewind"
	seekTimeStr := queryParams.Get("seek") // Time adjustment in seconds

	// Handle the command
	switch action {
	case "play":
		fmt.Fprintln(w, "Play command received.")
	case "pause":
		fmt.Fprintln(w, "Pause command received.")
	case "forward":
		// Logic for fast forwarding, e.g., "seek=10" to skip 10 seconds ahead
		fmt.Fprintf(w, "Fast-forward by %s seconds.\n", seekTimeStr)
	case "rewind":
		// Logic for rewinding, e.g., "seek=5" to skip back 5 seconds
		fmt.Fprintf(w, "Rewind by %s seconds.\n", seekTimeStr)
	default:
		http.Error(w, "Invalid action", http.StatusBadRequest)
		return
	}
}

func main() {
	http.HandleFunc("/video", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
