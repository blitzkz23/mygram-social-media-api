package rest

import (
	"mygram-social-media-api/dto"
	"mygram-social-media-api/entity"
	"mygram-social-media-api/pkg/helpers"
	"mygram-social-media-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type commentRestHandler struct {
	commentService service.CommentService
}

func NewCommentRestHandler(commentService service.CommentService) *commentRestHandler {
	return &commentRestHandler{commentService: commentService}
}

func (c *commentRestHandler) PostComment(ctx *gin.Context) {
	var commentRequest dto.CommentRequest
	var err error
	var userData entity.User

	contentType := helpers.GetContentType(ctx)
	if contentType == helpers.AppJSON {
		err = ctx.ShouldBindJSON(&commentRequest)
	} else {
		err = ctx.ShouldBind(&commentRequest)
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad_request",
			"message": err.Error(),
		})
		return
	}

	if value, ok := ctx.MustGet("userData").(entity.User); !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"err_message": "unauthorized",
		})
		return
	} else {
		userData = value
	}

	comment, err2 := c.commentService.PostComment(userData.ID, &commentRequest)
	if err2 != nil {
		ctx.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, comment)
}

func (c *commentRestHandler) GetAllComments(ctx *gin.Context) {
	var userData entity.User
	if value, ok := ctx.MustGet("userData").(entity.User); !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"err_message": "unauthorized",
		})
		return
	} else {
		userData = value
	}
	_ = userData

	comments, err := c.commentService.GetAllComments()
	if err != nil {
		ctx.JSON(err.Status(), gin.H{
			"error":   err.Error(),
			"message": err.Message(),
		})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}
