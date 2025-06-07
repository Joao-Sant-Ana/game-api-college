package handlers

import (
	"c02-project/internals/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		db: db,
	}
}

func (u *UserHandler) GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.User
		if result := u.db.Limit(20).Find(&users); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Database error",
			})
			return
		}

		if len(users) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "No users recorded yet",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	}
}

func (u *UserHandler) VerifyName() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		if name == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Please insert a username",
			})
			return
		}

		var user models.User
		err := u.db.Where("name = ?", name).First(&user).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNoContent, nil)
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Database error",
			})
			return
		}

		c.JSON(http.StatusConflict, gin.H{
			"message": "Name already in use",
		})
	}
}

func (u *UserHandler) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid body",
			})
			return
		}

		if result := u.db.Create(&user); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Database error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User created",
		})
	}
}

func (u *UserHandler) UpdateWaves() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Name string `json:"name"`
			Wave string `json:"wave"`
		}
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid body",
			})
			return
		}

		result := u.db.Model(&models.User{}).Where("name = ?", input.Name).Update("wave", input.Wave)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Database error",
			})
			return
		}
		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "User not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User updated",
		})
	}
}
