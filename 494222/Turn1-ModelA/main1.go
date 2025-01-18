package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Data struct {
	Message string `json:"message"`
}

func readFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
	}

	fileData, err := ioutil.ReadFile("data.json")
	if err != nil {
		http.Error(w, "Could not read file", http.StatusInternalServerError)
		return
	}

	var data Data
	err = json.Unmarshal(fileData, &data)
	if err != nil {
		http.Error(w, "Could not unmarshal data", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/read", readFileHandler)
	http.ListenAndServe(":8081", nil)
}
