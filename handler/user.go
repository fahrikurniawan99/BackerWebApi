package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "")

	response := helper.APIResponse("Accout has ben register", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	{
		var input user.LoginUserInput

		err := c.ShouldBindJSON(&input)

		if err != nil {
			errors := helper.FormatValidationError(err)

			errorMessage := gin.H{"errors": errors}

			response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		loggedUser, err := h.userService.Login(input)

		if err != nil {

			errorMessage := gin.H{"errors": err.Error()}

			response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		formatter := user.FormatUser(loggedUser, "")

		response := helper.APIResponse("Successfuly loggedin", http.StatusOK, "success", formatter)

		c.JSON(http.StatusOK, response)

	}
}

func (h *userHandler) CheckEmailAvaibality(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)

	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{"is_avalaible": isEmailAvailable}

	metaMessage := "Email is has been registered"

	if isEmailAvailable {
		metaMessage = "Email is avalaible"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)

}
