package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

const (
	filePath = "tmp/data.txt" // Replace with your desired file path
)

func readFileHandler(w http.ResponseWriter, r *http.Request) {
	// Read data from the file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Write the data to the response body
	w.Write(data)
}

func main() {
	http.HandleFunc("/read", readFileHandler)
	log.Println("Starting file reader service...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
