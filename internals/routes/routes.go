package routes

import (
	"c02-project/internals/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(r *gin.Engine, db *gorm.DB) {
	userHandler := handlers.NewUserHandler(db)

	r.GET("/users", userHandler.GetUsers())
	r.GET("/user/:name", userHandler.VerifyName())
	r.POST("/user", userHandler.CreateUser())
	r.PUT("/user", userHandler.UpdateWaves())
}
