package handler

import (
	"agolang/project-3/database"
	"agolang/project-3/repository/user_repository/user_pg"
	"agolang/project-3/service"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	var port = "8080"
	database.InitiliazeDatabase()

	db := database.GetDatabaseInstance()

	userRepo := user_pg.NewUserPG(db)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	authService := service.NewAuthService(userRepo)

	route := gin.Default()

	userRoute := route.Group("/users")
	{
		userRoute.POST("/register", userHandler.Register)

		userRoute.POST("/login", userHandler.Login)

		userRoute.PUT("/update-account", authService.Authentication(), userHandler.UpdateUser)
	}

	route.Run(":" + port)
}
