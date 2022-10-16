package rest

import (
	"fmt"
	"mygram-social-media-api/dto"
	"mygram-social-media-api/helpers"
	"mygram-social-media-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	appJSON = "application/json"
)

type userRestHandler struct {
	userService service.UserService
}

func NewUserRestHandler(userService service.UserService) *userRestHandler {
	return &userRestHandler{userService: userService}
}

func (u *userRestHandler) Login(c *gin.Context) {
	var user dto.LoginRequest
	var err error

	contentType := helpers.GetContentType(c)
	if contentType == appJSON {
		err = c.ShouldBindJSON(&user)
	} else {
		err = c.ShouldBind(&user)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad_request",
			"message": err.Error(),
		})
		return
	}

	token, err := u.userService.Login(&user)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "unprocessable_entity",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, token)
}

func (u *userRestHandler) Register(c *gin.Context) {
	var user dto.RegisterRequest
	var err error

	contentType := helpers.GetContentType(c)
	if contentType == appJSON {
		// ! TODO: JSON bind not working
		err = c.ShouldBindJSON(&user)
	} else {
		err = c.ShouldBind(&user)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad_request",
			"message": err.Error(),
		})
		return
	}

	fmt.Println("User =>", &user)
	successMessage, err := u.userService.Register(&user)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "unprocessable_entity",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, successMessage)
}
