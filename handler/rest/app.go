package rest

import (
	"fmt"
	"mygram-social-media-api/database"
	"mygram-social-media-api/repository/comment_repository/comment_pg"
	"mygram-social-media-api/repository/photo_repository/photo_pg"
	"mygram-social-media-api/repository/social_media_repository/social_media_pg"
	"mygram-social-media-api/repository/user_repository/user_pg"
	"mygram-social-media-api/service"

	"github.com/gin-gonic/gin"
)

const (
	port = "127.0.0.1:8080"
)

func StartApp() {
	database.StartDB()

	// ! Inject all the dependencies here
	db := database.GetDB()
	userRepo := user_pg.NewUserPG(db)
	userService := service.NewUserService(userRepo)
	userRestHandler := NewUserRestHandler(userService)

	photoRepo := photo_pg.NewPhotoPG(db)
	photoService := service.NewPhotoService(photoRepo)
	photoRestHandler := NewPhotoRestHandler(photoService)

	commentRepo := comment_pg.NewCommentPG(db)
	commentService := service.NewCommentService(commentRepo)
	commentRestHandler := NewCommentRestHandler(commentService)

	socialMediaRepo := social_media_pg.NewSocialMediaPG(db)
	socialMediaService := service.NewSocialMediaService(socialMediaRepo)
	socialMediaRestHandler := NewSocialMediaRestHandler(socialMediaService)

	authService := service.NewAuthService(userRepo, photoRepo, commentRepo, socialMediaRepo)

	// ! Routing
	route := gin.Default()

	userRoute := route.Group("/users")
	{
		userRoute.POST("/login", userRestHandler.Login)
		userRoute.POST("/register", userRestHandler.Register)
		userRoute.PUT("/", authService.Authentication(), userRestHandler.UpdateUserData)
		userRoute.DELETE("/", authService.Authentication(), userRestHandler.DeleteUser)
	}

	photoRoute := route.Group("/photos")
	{
		photoRoute.Use(authService.Authentication())
		photoRoute.POST("/", photoRestHandler.PostPhoto)
		photoRoute.GET("/", photoRestHandler.GetAllPhotos)
		photoRoute.PUT("/:photoID", authService.PhotoAuthorization(), photoRestHandler.UpdatePhoto)
		photoRoute.DELETE("/:photoID", authService.PhotoAuthorization(), photoRestHandler.DeletePhoto)
	}

	commentRoute := route.Group("/comments")
	{
		commentRoute.Use(authService.Authentication())
		commentRoute.POST("/", commentRestHandler.PostComment)
		commentRoute.GET("/", commentRestHandler.GetAllComments)
		commentRoute.PUT("/:commentID", authService.CommentAuthorization(), commentRestHandler.UpdateComment)
		commentRoute.DELETE("/:commentID", authService.CommentAuthorization(), commentRestHandler.DeleteComment)
	}

	socialMediaRoute := route.Group("/socialmedias")
	{
		socialMediaRoute.Use(authService.Authentication())
		socialMediaRoute.POST("/", socialMediaRestHandler.AddSocialMedia)
		socialMediaRoute.GET("/", socialMediaRestHandler.GetAllSocialMedias)
		socialMediaRoute.PUT("/:socialMediaID", authService.SocialMediaAuthorization(), socialMediaRestHandler.EditSocialMediaData)
		socialMediaRoute.DELETE("/:socialMediaID", authService.SocialMediaAuthorization(), socialMediaRestHandler.DeleteSocialMedia)
	}

	fmt.Println("Server running on PORT =>", port)
	route.Run(port)
}
