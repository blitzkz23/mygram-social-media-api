package rest

import (
	"fmt"
	"mygram-social-media-api/docs"
	"mygram-social-media-api/src/database"
	"mygram-social-media-api/src/repository/comment_repository/comment_pg"
	"mygram-social-media-api/src/repository/photo_repository/photo_pg"
	"mygram-social-media-api/src/repository/social_media_repository/social_media_pg"
	"mygram-social-media-api/src/repository/user_repository/user_pg"
	"mygram-social-media-api/src/service"
	"os"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var (
	PORT = os.Getenv("PORT")
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
	commentService := service.NewCommentService(commentRepo, photoRepo)
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

	// ! Docs
	docs.SwaggerInfo.Title = "MyGram API"
	docs.SwaggerInfo.Description = "MyGram is a social media API that allows users to post photos, comments, and add social media links. Created with Golang, Gin-gonic, GORM, and PostgreSQL utilizing DDD pattern."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "https://mygram-social-media-api-production.up.railway.app/"
	// docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"https"}

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	fmt.Println("Server running on PORT =>", PORT)
	route.Run(":" + PORT)
}
