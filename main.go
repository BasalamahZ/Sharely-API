package main

import (
	"net/http"
	"sharely/configs"
	"sharely/controllers"
	"sharely/middlewares"
	"sharely/repositories"
	"sharely/services"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.LoadEnvVariables()
	configs.ConnectDB()
	configs.SyncDB()
}

func main() {
	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "welcome home",
		})
	})
	authRepositry := repositories.NewAuthRepository()
	authService := services.NewAuthService(&authRepositry)
	authController := controllers.NewAuthController(&authService)
	server.POST("/register", authController.Register)
	server.POST("/login", authController.Login)
	server.GET("/user", middlewares.VerifyAuth, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message" : "authorization",
		})
	})
	server.Run()
}