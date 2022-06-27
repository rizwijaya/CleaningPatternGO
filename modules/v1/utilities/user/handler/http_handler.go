package user

import (
	"ClearningPatternGO/app/helper"
	"ClearningPatternGO/modules/v1/utilities/user/models"
	ss "ClearningPatternGO/modules/v1/utilities/user/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	RegisterUser(c *gin.Context)
}

type userHandler struct {
	userService ss.Service
}

func NewUserHandler(userService ss.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input models.RegisterUserInput

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

	formatter := models.FormatUser(newUser, token)

	response := helper.APIRespon("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
