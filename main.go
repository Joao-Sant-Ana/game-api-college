package main

import (
	"c02-project/config"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	router := config.StartGin()
	db := config.ConnectDB()
}
