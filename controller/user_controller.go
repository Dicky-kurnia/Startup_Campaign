package controller

import (
	"bwastartup/formatter"
	"bwastartup/helper"
	"bwastartup/model"
	"bwastartup/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.ServiceUser
}

func NewUserController(userService service.ServiceUser) *userController {
	return &userController{userService}
}

func (controler *userController) RegisterUserController(c *gin.Context) {
	var input model.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newuser, err := controler.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "Succsess", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := formatter.FormatUser(newuser, "tokentokentoken")
	response := helper.APIResponse("Account has been registered", http.StatusOK, "Succsess", formatter)
	c.JSON(http.StatusOK, response)

}

func (controller *userController) Login(c *gin.Context) {

}
