package rest

import (
	"fmt"
	"mygram-social-media-api/dto"
	"mygram-social-media-api/entity"
	"mygram-social-media-api/pkg/helpers"
	"mygram-social-media-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userRestHandler struct {
	userService service.UserService
}

func NewUserRestHandler(userService service.UserService) *userRestHandler {
	return &userRestHandler{userService: userService}
}

func (u *userRestHandler) Login(c *gin.Context) {
	var userRequest dto.LoginRequest
	var err error

	contentType := helpers.GetContentType(c)
	if contentType == helpers.AppJSON {
		err = c.ShouldBindJSON(&userRequest)
	} else {
		err = c.ShouldBind(&userRequest)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad_request",
			"message": err.Error(),
		})
		return
	}

	token, err := u.userService.Login(&userRequest)

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
	var userRequest dto.RegisterRequest
	var err error

	contentType := helpers.GetContentType(c)
	if contentType == helpers.AppJSON {
		// ! TODO: JSON bind not working
		err = c.ShouldBindJSON(&userRequest)
	} else {
		err = c.ShouldBind(&userRequest)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad_request",
			"message": err.Error(),
		})
		return
	}

	fmt.Println("User =>", &userRequest)
	successMessage, err := u.userService.Register(&userRequest)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "unprocessable_entity",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, successMessage)
}

func (u *userRestHandler) UpdateUserData(c *gin.Context) {
	var updateUserDataRequest dto.UpdateUserDataRequest
	var err error
	var userData entity.User

	contentType := helpers.GetContentType(c)
	if contentType == helpers.AppJSON {
		err = c.ShouldBindJSON(&updateUserDataRequest)
	} else {
		err = c.ShouldBind(&updateUserDataRequest)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad_request",
			"message": err.Error(),
		})
		return
	}

	if value, ok := c.MustGet("userData").(entity.User); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err_message": "unauthorized",
		})
		return
	} else {
		userData = value
	}
	fmt.Println("APAKAH ADA ID", userData.ID)

	// ! TODO: Update error but data updated
	response, err := u.userService.UpdateUserData(userData.ID, &updateUserDataRequest)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "unprocessable_entity",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (u *userRestHandler) DeleteUser(c *gin.Context) {
	var userData entity.User
	if value, ok := c.MustGet("userData").(entity.User); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err_message": "unauthorized",
		})
		return
	} else {
		userData = value
	}
	fmt.Println("APAKAH ADA ID", userData.ID)

	response, err := u.userService.DeleteUser(userData.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal_server_error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
