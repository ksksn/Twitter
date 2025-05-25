package main

import (
	"log"
	routes "twitter/server"
	"twitter/db"
)

func main() {
	log.Println("Starting Twitter server...")
	if err := db.InitDB(); err != nil {
		log.Fatalf("Не удалось инициализировать базу данных: %v", err)
	}
	if err := db.InitSchema(); err != nil {
		log.Fatalf("Не удалось инициализировать базу данных: %v", err)
	}
	defer db.CloseDB()
	router := routes.NewRouter()
	router.SetupRoutes()

	log.Println("Server listening on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	
}
