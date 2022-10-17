package service

import (
	"fmt"
	"mygram-social-media-api/entity"
	"mygram-social-media-api/pkg/helpers"
	"mygram-social-media-api/repository/photo_repository"
	"mygram-social-media-api/repository/user_repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	PhotoAuthorization() gin.HandlerFunc
}

type authService struct {
	userRepository  user_repository.UserRepository
	photoRepository photo_repository.PhotoRepository
}

func NewAuthService(userRepository user_repository.UserRepository, photoRepository photo_repository.PhotoRepository) AuthService {
	return &authService{
		userRepository:  userRepository,
		photoRepository: photoRepository,
	}
}

func (a *authService) Authentication() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var user *entity.User = &entity.User{}

		// Get token from header
		tokenStr := ctx.Request.Header.Get("Authorization")
		err := user.VerifyToken(tokenStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error_message": err.Error(),
			})
			return
		}

		_, err = a.userRepository.GetUserByIDAndEmail(user)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error_message": err.Error(),
			})
			return
		}

		ctx.Set("userData", *user)
		ctx.Next()
	})
}

func (a *authService) PhotoAuthorization() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var userData entity.User

		if value, ok := ctx.MustGet("userData").(entity.User); !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error_message": "unauthorized",
			})
			return
		} else {
			userData = value
		}

		photoIdParam, err := helpers.GetParamId(ctx, "photoID")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"err_message": "invalid params",
			})
			return
		}
		fmt.Println("Apakah ada ID", photoIdParam)

		photo, err := a.photoRepository.GetPhotoByID(photoIdParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"err_message": "photo not found",
			})
			return
		}
		fmt.Println("Apakah ada photo", &photo)

		if photo.UserID != userData.ID {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"err_message": "forbidden access",
			})
			return
		}

		ctx.Next()
	})
}
