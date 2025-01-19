package main

import (
	"fmt"
	"net/http"
)

func videoHandler(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()

	playbackSpeed := queries.Get("speed")
	resolution := queries.Get("resolution")
	tutorial := queries.Get("tutorial")

	// Default settings if parameters are not provided.
	if playbackSpeed == "" {
		playbackSpeed = "1.0"
	}
	if resolution == "" {
		resolution = "1080p"
	}
	if tutorial == "" {
		tutorial = "basic"
	}

	// Use these parameters to modify video playback or load tutorials.
	fmt.Fprintf(w, "Playback Speed: %s\nResolution: %s\nTutorial: %s\n", playbackSpeed, resolution, tutorial)
}

func main() {
	http.HandleFunc("/video", videoHandler)
	http.ListenAndServe(":8080", nil)
}
