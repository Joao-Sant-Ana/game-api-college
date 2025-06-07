package routes

import (
	"c02-project/internals/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	userHandler := handlers.NewUserHandler(db)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/users", userHandler.GetUsers())
	r.GET("/user/:name", userHandler.VerifyName())
	r.POST("/user", userHandler.CreateUser())
	r.PATCH("/user", userHandler.UpdateWaves())
}
