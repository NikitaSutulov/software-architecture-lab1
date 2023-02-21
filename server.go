package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func timeHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/time" {
		http.Error(responseWriter, "404 not found", http.StatusNotFound)
		return
	}
	if request.Method != http.MethodGet {
		http.Error(responseWriter, "Method is not supported", http.StatusNotFound)
		return
	}
	currentTime := time.Now().Format(time.RFC3339)
	response := map[string]string{"time": currentTime}
	responseWriter.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(responseWriter).Encode(response)
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

// i play pokemon go every day
