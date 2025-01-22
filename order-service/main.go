package main

import (
	"log"
	"net/http"
)

func main() {
	router, err := InitializeRouter()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	log.Println("Order service is running on port 8003...")
	log.Fatal(http.ListenAndServe(":8003", router))
}
