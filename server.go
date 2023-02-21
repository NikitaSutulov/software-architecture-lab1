package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/time" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
		}
	if r.Method != http.MethodGet {
	http.Error(w, "Method is not supported", http.StatusNotFound)
	return
	}

	currentTime := time.Now().Format(time.RFC3339)
	response := map[string]string{"time": currentTime}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
	return
	}
}

func main() {
	http.HandleFunc("/time", timeHandler)

	fmt.Printf("Starting server at port 8795\n")
	if err := http.ListenAndServe(":8795", nil); err != nil {
		log.Fatal(err)
	}
}
