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

type userUpdate struct {
	Name string `json:"name" example:"joao"`
	Wave int `json:"wave" example:"10"`
}

type userCreate struct {
	Name string `json:"name" example:"joao"`
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		db: db,
	}
}

// GetUsers godoc
// @Summary      Get list of users
// @Description  Returns up to 20 users
// @Tags         users
// @Produce      json
// @Success      200  {object}  []models.User
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /users [get]
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

// VerifyName godoc
// @Summary      Check if username is taken
// @Description  Verifies if a username already exists
// @Tags         users
// @Param        name   path      string  true  "Username to check"
// @Success      204    "No Content, name available"
// @Failure      400    {object}  map[string]string
// @Failure      409    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /users/verify/{name} [get]
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

// CreateUser godoc
// @Summary      Create a new user
// @Description  Creates a user with the given JSON body
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      userCreate  true  "User to create"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /users [post]
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

// UpdateWaves godoc
// @Summary      Update user's wave field
// @Description  Updates the wave field for a user by name
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        input  body      userUpdate  true  "Update payload"
// @Success      200    {object}  map[string]string
// @Failure      400    {object}  map[string]string
// @Failure      404    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /users/waves [patch]
func (u *UserHandler) UpdateWaves() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Name string `json:"name"`
			Wave int `json:"wave"`
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
