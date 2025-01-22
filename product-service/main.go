package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize the router with dependencies using Wire
	router, err := InitializeRouter()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// Start server
	log.Println("Product service is running on port 8001...")
	log.Fatal(http.ListenAndServe(":8001", router))
}
