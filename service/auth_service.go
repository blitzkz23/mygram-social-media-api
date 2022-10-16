package service

import (
	"mygram-social-media-api/entity"
	"mygram-social-media-api/repository/user_repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
}

type authService struct {
	userRepository user_repository.UserRepository
}

func NewAuthService(userRepository user_repository.UserRepository) AuthService {
	return &authService{userRepository: userRepository}
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
