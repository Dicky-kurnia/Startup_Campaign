package main

import (
	"bwastartup/config"
	"bwastartup/controller"
	"bwastartup/middleware"
	"bwastartup/repository"
	"bwastartup/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	mysql, err := config.MySQLConnection()
	if err != nil {
		fmt.Println("Failed to database connection:", err)
		return
	}

	defer func() {
		if sqlDB, err := mysql.DB(); err == nil {
			sqlDB.Close()
		}
	}()

	userRepository := repository.NewUserRepository(mysql)
	userService := service.NewServiceUser(userRepository)
	authService := middleware.NewAuthService()
	userController := controller.NewUserController(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	//user
	api.POST("/users", userController.RegisterUserController)
	api.POST("/sessions", userController.Login)
	api.POST("/email_checkers", userController.CheckEmailAvailability)
	api.POST("/avatars", userController.UploadAvatar)

	router.Run()
}
