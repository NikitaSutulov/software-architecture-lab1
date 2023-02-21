package main

import (
	"fmt"
	"log"
	"net/http"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func main() {
	http.HandleFunc("/time", timeHandler)

	fmt.Printf("Starting server at port 8795\n")
	if err := http.ListenAndServe(":8795", nil); err != nil {
		log.Fatal(err)
	}
}
