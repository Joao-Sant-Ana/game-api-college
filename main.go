package main

import (
	"c02-project/config"
	"c02-project/internals/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	router := config.StartGin()
	db := config.ConnectDB()

	routes.SetupRoutes(router, db)
}
