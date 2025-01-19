package main

import (
	"fmt"
	"log"
	"net/http"
)

func playbackControlHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the URL query parameters
	query := r.URL.Query()

	// Check if the "action" parameter is present
	action := query.Get("action")
	if action == "" {
		http.Error(w, "Invalid playback control action", http.StatusBadRequest)
		return
	}

	// Process the playback control action based on its value
	switch action {
	case "play":
		fmt.Fprintf(w, "Playing the video...")
	case "pause":
		fmt.Fprintf(w, "Pausing the video...")
	case "fastforward":
		fmt.Fprintf(w, "Fast-forwarding the video...")
	case "rewind":
		fmt.Fprintf(w, "Rewinding the video...")
	default:
		http.Error(w, "Invalid playback control action", http.StatusBadRequest)
		return
	}
}

func main() {
	// Set up the server and handle requests
	http.HandleFunc("/playback", playbackControlHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
