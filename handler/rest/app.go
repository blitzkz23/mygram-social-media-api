package rest

import (
	"fmt"
	"mygram-social-media-api/database"
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

	authService := service.NewAuthService(userRepo)

	// ! Routing
	route := gin.Default()

	userRoute := route.Group("/user")
	{
		userRoute.POST("/login", userRestHandler.Login)
		userRoute.POST("/register", userRestHandler.Register)
		userRoute.PUT("/update/:userID", authService.Authentication(), userRestHandler.UpdateUserData)
		userRoute.DELETE("/delete/:userID", authService.Authentication(), userRestHandler.DeleteUser)
	}

	fmt.Println("Server running on PORT =>", port)
	route.Run(port)
}
