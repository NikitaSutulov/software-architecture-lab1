package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// function timeHandler
// handles GET request to route /time
// sends current time in RFC3339 format in JSON-response
func timeHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/time" {
		http.Error(responseWriter, "404 not found", http.StatusNotFound)
		return
	}
	if request.Method != http.MethodGet {
		http.Error(responseWriter, "Method is not supported", http.StatusNotFound)
		return
	}
	// Retrieves current time in RFC3339 format, assigning it to the variable "currentTime".
	currentTime := time.Now().Format(time.RFC3339)
	// Creates a map variable response with a key "time" and the value currentTime.
	response := map[string]string{"time": currentTime}
	// Sets the response header "Content-Type" to "application/json".
	responseWriter.Header().Set("Content-Type", "application/json")
	// Encodes the response variable as JSON and writes it to the response writer, returning any errors encountered.
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
