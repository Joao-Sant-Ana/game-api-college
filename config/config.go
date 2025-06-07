package config

import (
	"c02-project/internals/models"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func readEnv(key, defaultValue string) string {
	if value, has := os.LookupEnv(key); has {
		return value
	}

	return defaultValue
}

func StartGin() *gin.Engine {
	router := gin.Default()

	return router
}

func ConnectDB() *gorm.DB {
	dsn := readEnv("DB_URL", "")
	if dsn == "" {
		log.Fatalf("No database URL provided")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)

	db.AutoMigrate(&models.User{})

	return db
}
