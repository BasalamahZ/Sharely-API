package main

import (
	"sharely/configs"
	"sharely/repositories"
	"sharely/services"
	"sharely/controllers"
	

	"github.com/gin-gonic/gin"
)

func init() {
	configs.LoadEnvVariables()
	configs.ConnectDB()
	configs.SyncDB()
}

func main() {
	server := gin.Default()
	authRepositry := repositories.NewAuthRepository()
	authService := services.NewAuthService(&authRepositry)
	authController := controllers.NewAuthController(&authService)
	server.POST("/register", authController.Register)
	server.POST("/login", authController.Login)
	server.Run()
}
