package controller

import (
	"bwastartup/formatter"
	"bwastartup/helper"
	"bwastartup/middleware"
	"bwastartup/model"
	"bwastartup/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.ServiceUser
	authService middleware.AuthService
}

func NewUserController(userService service.ServiceUser, authService middleware.AuthService) *userController {
	return &userController{userService, authService}
}

func (controller *userController) RegisterUserController(c *gin.Context) {
	var input model.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newuser, err := controller.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := controller.authService.GenerateToken(newuser.ID)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := formatter.FormatUser(newuser, token)

	response := helper.APIResponse("Account has been registered", http.StatusOK, "Succsess", formatter)
	c.JSON(http.StatusOK, response)

}

func (controller *userController) Login(c *gin.Context) {
	var input model.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := controller.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := controller.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := formatter.FormatUser(loggedinUser, token)

	response := helper.APIResponse("Succsess fully loggedin", http.StatusOK, "Succsess", formatter)
	c.JSON(http.StatusOK, response)

}

func (controller *userController) CheckEmailAvailability(c *gin.Context) {
	var input model.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := controller.userService.IsEmailEvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server Error"}

		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has ben registered"

	if isEmailAvailable {
		metaMessage = "Email is Available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "succsess", data)
	c.JSON(http.StatusOK, response)
}

func (controller *userController) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("failed to upload avatar image", http.StatusBadRequest, "Error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID := 1
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("failed to upload avatar image", http.StatusBadRequest, "Error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = controller.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("failed to upload avatar image", http.StatusBadRequest, "Error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": false}
	response := helper.APIResponse("avatar succsess fully uploaded", http.StatusOK, "Succsess", data)
	c.JSON(http.StatusOK, response)

}
