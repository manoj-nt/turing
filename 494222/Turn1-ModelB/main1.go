package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	filePath = "tmp/data.txt" // Replace with your desired file path
)

func writeFileHandler(w http.ResponseWriter, r *http.Request) {
	// Read the data from the request body
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Write the data to the file
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		log.Printf("Error writing file: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Write a success response
	fmt.Fprintf(w, "Data written to file successfully!")
}

func main() {
	http.HandleFunc("/write", writeFileHandler)
	log.Println("Starting file writer service...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
