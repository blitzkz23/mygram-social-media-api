package rest

import (
	"mygram-social-media-api/dto"
	"mygram-social-media-api/entity"
	"mygram-social-media-api/pkg/helpers"
	"mygram-social-media-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type photoRestHandler struct {
	photoService service.PhotoService
}

func NewPhotoRestHandler(photoService service.PhotoService) *photoRestHandler {
	return &photoRestHandler{photoService: photoService}
}

func (p *photoRestHandler) PostPhoto(c *gin.Context) {
	var photoRequest dto.PhotoRequest
	var err error
	var userData entity.User

	contentType := helpers.GetContentType(c)
	if contentType == helpers.AppJSON {
		err = c.ShouldBindJSON(&photoRequest)
	} else {
		err = c.ShouldBind(&photoRequest)
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

	photo, err := p.photoService.PostPhoto(userData.ID, &photoRequest)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "unprocessable_entity",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, photo)
}

func (p *photoRestHandler) GetAllPhotos(c *gin.Context) {
	var userData entity.User
	if value, ok := c.MustGet("userData").(entity.User); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err_message": "unauthorized",
		})
		return
	} else {
		userData = value
	}
	_ = userData

	photos, err := p.photoService.GetAllPhotos()
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"err_message": "forbidden",
		})
		return
	}

	c.JSON(http.StatusOK, photos)

}
