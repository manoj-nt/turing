package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Data struct {
	Message string `json:"message"`
}

func writeFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	file, err := os.Create("data.json")
	if err != nil {
		http.Error(w, "Could not create file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Could not marshal data", http.StatusInternalServerError)
		return
	}

	_, err = file.Write(jsonData)
	if err != nil {
		http.Error(w, "Could not write to file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Data written to file successfully")
}

func main() {
	http.HandleFunc("/write", writeFileHandler)
	http.ListenAndServe(":8080", nil)
}
