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

	log.Println("User service is running on port 8002...")
	log.Fatal(http.ListenAndServe(":8002", router))
}
