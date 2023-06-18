package handler

import (
	"agolang/project-3/database"
	"agolang/project-3/repository/category_repository/category_pg"
	"agolang/project-3/repository/task_repository/task_pg"
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

	categoryRepo := category_pg.NewCategoryPG(db)
	taskRepo := task_pg.NewTaskPG(db)

	categoryService := service.NewCategoryService(categoryRepo, taskRepo)
	taskService := service.NewTaskService(taskRepo, categoryRepo, userRepo)

	categoryHandler := NewCategoryHandler(categoryService)
	taskHandler := NewTaskHandler(taskService)

	_ = taskHandler

	authService := service.NewAuthService(userRepo, taskRepo)

	route := gin.Default()

	userRoute := route.Group("/users")
	{
		userRoute.POST("/register", userHandler.Register)

		userRoute.POST("/login", userHandler.Login)

		userRoute.PUT("/update-account", authService.Authentication(), userHandler.UpdateUser)

		userRoute.DELETE("/delete-account", authService.Authentication(), userHandler.DeleteUser)

	}

	categoryRoute := route.Group("/categories")
	{
		categoryRoute.POST("/", authService.Authentication(), authService.AdminAuthorization(), categoryHandler.CreateCategory)

		categoryRoute.GET("/", authService.Authentication(), categoryHandler.GetAllCategories)

		categoryRoute.PATCH("/:categoryId", authService.Authentication(), authService.AdminAuthorization(), categoryHandler.UpdateCategory)

		categoryRoute.DELETE("/:categoryId", authService.Authentication(), authService.AdminAuthorization(), categoryHandler.DeleteCategory)

	}

	taskRoute := route.Group("/tasks")
	{
		taskRoute.POST("/", authService.Authentication(), taskHandler.CreateTask)

		taskRoute.GET("/", authService.Authentication(), taskHandler.GetAllTasks)

		taskRoute.PUT("/:taskId", authService.Authentication(), authService.TaskAuthorization(), taskHandler.UpdateTask)

		taskRoute.PATCH("/update-status/:taskId", authService.Authentication(), authService.TaskAuthorization(), taskHandler.UpdateTaskStatus)

		taskRoute.PATCH("/update-category/:taskId", authService.Authentication(), authService.TaskAuthorization(), taskHandler.UpdateTaskCategory)

		taskRoute.DELETE("/:taskId", authService.Authentication(), authService.TaskAuthorization(), taskHandler.DeleteTask)
	}

	route.Run(":" + port)
}
