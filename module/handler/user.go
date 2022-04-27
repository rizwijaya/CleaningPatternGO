package handler

import (
	"TamaskaPJU/app/helper"
	"TamaskaPJU/module/utilities/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	RegisterUser(c *gin.Context)
}

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

		response := helper.APIRespon("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIRespon("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	token := "JWTJWTWJJWTJWJW"
	// token, err := h.authService.GenerateToken(newUser.ID)
	// if err != nil {
	// 	response := helper.APIRespon("Register account failed", http.StatusBadRequest, "error", nil)
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	formatter := user.FormatUser(newUser, token)

	response := helper.APIRespon("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
