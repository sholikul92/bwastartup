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
	var input user.InputRegistUser

	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.ErrorValidation(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Registered account failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	users, errUsers := h.userService.Register(input)
	if errUsers != nil {
		response := helper.APIResponse("Registered account failed", http.StatusInternalServerError, "failed", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	userFormatter := user.UserFormatter(*users, "token123")

	response := helper.APIResponse("Your account has been registed", http.StatusOK, "string", userFormatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.InputLoginUser

	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.ErrorValidation(err)

		errMsg := gin.H{"errors": errors}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "failed", errMsg)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	users, err := h.userService.Login(input)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusInternalServerError, "failed", err.Error())

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	userFormatter := user.UserFormatter(users, "tokentoken")

	response := helper.APIResponse("Login success", http.StatusOK, "success", userFormatter)
	c.JSON(http.StatusOK, response)
}
