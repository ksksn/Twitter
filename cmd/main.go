package main

import (
	"log"
	routes "twitter/server"
)

func main() {
	log.Println("Starting Twitter server...")
	router := routes.NewRouter()
	router.SetupRoutes()

	log.Println("Server listening on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
