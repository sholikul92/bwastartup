package handler

import (
	"bwastartup/auth"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	jwtService  auth.Service
}

func NewUserHandler(userService user.Service, jwtSevice auth.Service) *userHandler {
	return &userHandler{userService, jwtSevice}
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

	token, err := h.jwtService.GenerateToken(users.ID)
	if err != nil {
		response := helper.APIResponse("Registered account failed", http.StatusInternalServerError, "failed", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	userFormatter := user.UserFormatter(*users, token)

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

	token, err := h.jwtService.GenerateToken(users.ID)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusInternalServerError, "failed", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	userFormatter := user.UserFormatter(users, token)

	response := helper.APIResponse("Login success", http.StatusOK, "success", userFormatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckedEmail(c *gin.Context) {
	var input user.CheckedEmailInput

	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.ErrorValidation(err)

		errMsg := gin.H{"errors": errors}

		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errMsg)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	IsEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errMsg := gin.H{"errors": "Server error"}

		response := helper.APIResponse("Invalid input", http.StatusInternalServerError, "error", errMsg)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	message := "Email sudah digunakan"

	if IsEmailAvailable {
		message = "Email tersedia"
	}

	data := gin.H{
		"is_available": IsEmailAvailable,
	}

	response := helper.APIResponse(message, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UpdateAvatarImage(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}

		response := helper.APIResponse("Failed to avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.Users)
	id := currentUser.ID
	path := fmt.Sprintf("images/%d-%s", id, file.Filename)

	// save file dari input form
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}

		response := helper.APIResponse("Failed to avatar image", http.StatusBadRequest, "failed", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(id, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}

		response := helper.APIResponse("Failed to avatar image", http.StatusInternalServerError, "error", data)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	data := gin.H{
		"is_upload": true,
	}
	response := helper.APIResponse("Success upload avatar", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
