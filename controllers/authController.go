package controllers

import (
	"net/http"
	"sharely/middlewares"
	"sharely/models"
	"sharely/services"

	"github.com/gin-gonic/gin"
)

type authController struct {
	authService services.AuthService
}

func (ac *authController) Register(c *gin.Context) {
	requestRegister := new(models.User)
	err := c.BindJSON(&requestRegister)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "cannot read the body",
		})
		return
	}
	user, err := ac.authService.Register(*requestRegister)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to create user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": user,
	})
}

func (ac *authController) Login(c *gin.Context) {
	requestLogin := new(models.LoginRequest)
	err := c.BindJSON(&requestLogin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "cannot read the body",
		})
		return
	}
	user, err := ac.authService.Login(*requestLogin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to create user",
		})
		return
	}

	userID := user.ID
	token, err := middlewares.GenerateToken(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : false,
			"message": "Failed To Generate Token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": token,
	})
}

func NewAuthController(authService *services.AuthService) authController {
	return authController{
		authService: *authService,
	}
}
