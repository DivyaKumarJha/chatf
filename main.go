package main

import (
	"backend/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	r := routes.Router()

	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" 
	}

	log.Printf("Server starting on port %s...", port)
	err := http.ListenAndServe("0.0.0.0:"+port, r)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
