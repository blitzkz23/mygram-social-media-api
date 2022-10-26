package rest

import (
	"mygram-social-media-api/src/dto"
	"mygram-social-media-api/src/entity"
	"mygram-social-media-api/src/pkg/helpers"
	"mygram-social-media-api/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type socialMediaRestHandler struct {
	socialMediaService service.SocialMediaService
}

func NewSocialMediaRestHandler(socialMediaService service.SocialMediaService) *socialMediaRestHandler {
	return &socialMediaRestHandler{socialMediaService: socialMediaService}
}

// AddSocialMedia godoc
// @Summary Add social media data to user's account
// @Tags socialmedias
// @Description Add Social Media to your account
// @ID add-social-media
// @Accept  json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.SocialMediaRequest true "Add social media request body json"
// @Success 201 {object} dto.SocialMediaResponse
// @Router /socialmedias [post]
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

// GetAllSocialMedias godoc
// @Summary Get all social media datas
// @Tags socialmedias
// @Description Get all social medias
// @ID get-social-medias
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} dto.GetSocialMediaResponse
// @Router /socialmedias [get]
func (s *socialMediaRestHandler) GetAllSocialMedias(c *gin.Context) {
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

	socialMedia, err := s.socialMediaService.GetAllSocialMedias()
	if err != nil {
		c.JSON(err.Status(), gin.H{
			"error":   err.Error(),
			"message": err.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, socialMedia)
}

// EditSocialMediaData godoc
// @Summary Edit existing social media data
// @Tags socialmedias
// @Description Edit social media data
// @ID edit-social-media
// @Accept  json
// @Produce json
// @Param socialMediaID path uint true "social media's id"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.SocialMediaRequest true "Edit social media request body json"
// @Success 200 {object} dto.UpdateSocialMediaResponse
// @Router /socialmedias/{socialMediaID} [put]
func (s *socialMediaRestHandler) EditSocialMediaData(c *gin.Context) {
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
	_ = userData

	socialMediaIdParam, err := helpers.GetParamId(c, "socialMediaID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err_message": "invalid params",
		})
		return
	}

	socialMedia, err2 := s.socialMediaService.EditSocialMediaData(socialMediaIdParam, &socialMediaRequest)
	if err2 != nil {
		c.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, socialMedia)
}

// DeleteSocialMediaData godoc
// @Summary Delete existing social media data
// @Tags socialmedias
// @Description Delete social media data
// @ID delete-social-media
// @Produce json
// @Param socialMediaID path uint true "social media's id"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} dto.DeleteSocialMediaResponse
// @Router /socialmedias/{socialMediaID} [delete]
func (s *socialMediaRestHandler) DeleteSocialMedia(c *gin.Context) {
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

	socialMediaIdParam, err := helpers.GetParamId(c, "socialMediaID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err_message": "invalid params",
		})
		return
	}

	res, err2 := s.socialMediaService.DeleteSocialMedia(socialMediaIdParam)
	if err2 != nil {
		c.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
