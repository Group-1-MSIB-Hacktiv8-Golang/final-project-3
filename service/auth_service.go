package service

import (
	"agolang/project-3/entity"
	"agolang/project-3/pkg/errs"
	"agolang/project-3/repository/task_repository"
	"agolang/project-3/repository/user_repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	AdminAuthorization() gin.HandlerFunc
	TaskAuthorization() gin.HandlerFunc
}

type authService struct {
	userRepo user_repository.UserRepository
	taskRepo task_repository.TaskRepository
}

func NewAuthService(userRepo user_repository.UserRepository, taskRepo task_repository.TaskRepository) AuthService {
	return &authService{
		userRepo: userRepo,
		taskRepo: taskRepo,
	}
}

func (a *authService) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.GetHeader("Authorization")

		var user entity.User

		if err := user.ValidateToken(bearerToken); err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		result, err := a.userRepo.GetUserById(user.Id)
		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}
		//
		ctx.Set("userData", result)
		ctx.Next()
	}
}

func (a *authService) AdminAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, ok := ctx.MustGet("userData").(*entity.User)
		if !ok {
			newError := errs.NewBadRequest("Failed to get user data")
			ctx.AbortWithStatusJSON(newError.Status(), newError)
			return
		}
		if userData.Role != "admin" {
			newError := errs.NewUnauthorizedError("You're not authorized to access this endpoint")
			ctx.AbortWithStatusJSON(newError.Status(), newError)
			return
		}

		ctx.Next()
	}
}

func (a *authService) TaskAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, ok := ctx.MustGet("userData").(*entity.User)
		if !ok {
			newError := errs.NewBadRequest("Failed to get user data")
			ctx.AbortWithStatusJSON(newError.Status(), newError)
			return
		}

		taskId := ctx.Param("taskId")
		taskIdInt, err := strconv.ParseUint(taskId, 10, 32)
		if err != nil {
			newError := errs.NewBadRequest("Task id should be an unsigned integer")
			ctx.AbortWithStatusJSON(newError.Status(), newError)
			return
		}

		task, err2 := a.taskRepo.GetTaskById(int(taskIdInt))
		if err2 != nil {
			ctx.AbortWithStatusJSON(err2.Status(), err2)
			return
		}

		if task.UserId != userData.Id {
			newError := errs.NewUnauthorizedError("You're not authorized to modify this task")
			ctx.AbortWithStatusJSON(newError.Status(), newError)
			return
		}

		ctx.Next()
	}
}
