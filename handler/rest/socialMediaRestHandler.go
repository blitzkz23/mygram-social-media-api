package rest

import (
	"mygram-social-media-api/dto"
	"mygram-social-media-api/entity"
	"mygram-social-media-api/pkg/helpers"
	"mygram-social-media-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type socialMediaRestHandler struct {
	socialMediaService service.SocialMediaService
}

func NewSocialMediaRestHandler(socialMediaService service.SocialMediaService) *socialMediaRestHandler {
	return &socialMediaRestHandler{socialMediaService: socialMediaService}
}

func (s *socialMediaRestHandler) AddSocialMedia(c *gin.Context) {
	var socialMediaRequest dto.SocialMediaRequest
	var err error
	var userData entity.User

	contentType := helpers.GetContentType(c)
	if contentType == helpers.AppJSON {
		err = c.ShouldBindJSON(&socialMediaRequest)
	} else {
		err = c.ShouldBind(&socialMediaRequest)
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

	socialMedia, err2 := s.socialMediaService.AddSocialMedia(userData.ID, &socialMediaRequest)
	if err2 != nil {
		c.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}

	c.JSON(http.StatusCreated, socialMedia)

}
