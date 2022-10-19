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

// PostPhoto godoc
// @Summary Post a new photo
// @Tags photos
// @Description Post a new photo
// @ID post-photo
// @Accept  json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.PhotoRequest true "Add photo request body json"
// @Success 201 {object} dto.UpdatePhotoResponse
// @Router /photos [post]
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

// GetAllPhotos godoc
// @Summary Get all photos
// @Tags photos
// @Description Get all photos
// @ID get-all-photos
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} dto.GetPhotoResponse
// @Router /photos [get]
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
			"error":   "forbidden",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, photos)
}

// UpdatePhoto godoc
// @Summary Update existing photo data
// @Tags photos
// @Description Update photo data
// @ID update-photo
// @Accept  json
// @Produce json
// @Param photoID path uint true "photo's id"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.PhotoRequest true "Edit photo request body json"
// @Success 200 {object} dto.UpdatePhotoResponse
// @Router /photos/{photoID} [put]
func (p *photoRestHandler) UpdatePhoto(c *gin.Context) {
	var photoRequest dto.PhotoRequest
	var err error

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

	photoIdParam, err := helpers.GetParamId(c, "photoID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err_message": "invalid params",
		})
		return
	}

	photo, err2 := p.photoService.EditPhotoData(photoIdParam, &photoRequest)
	if err2 != nil {
		c.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, photo)
}

// DeletePhoto godoc
// @Summary Delete existing photo
// @Tags photos
// @Description Delete photo
// @ID delete-photo
// @Produce json
// @Param photoID path uint true "photo's id"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} dto.DeletePhotoResponse
// @Router /photos/{photoID} [delete]
func (p *photoRestHandler) DeletePhoto(c *gin.Context) {
	photoIdParam, err := helpers.GetParamId(c, "photoID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err_message": "invalid params",
		})
		return
	}

	_ = photoIdParam

	res, err2 := p.photoService.DeletePhoto(photoIdParam)
	if err2 != nil {
		c.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
