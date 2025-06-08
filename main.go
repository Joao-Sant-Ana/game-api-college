package main

import (
	"c02-project/config"
	"c02-project/internals/routes"
	"log"

	"github.com/joho/godotenv"
)

// @title           Game API
// @version         1.0
// @description    	Simple API for a game made in construct3
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	router := config.StartGin()
	db := config.ConnectDB()

	routes.SetupRoutes(router, db)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting API: %v", err)
	}
}
