package handler

import (
	"agolang/project-3/dto"
	"agolang/project-3/entity"
	"agolang/project-3/pkg/errs"
	"agolang/project-3/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{
		userService: userService,
	}
}

func (uh *userHandler) Register(ctx *gin.Context) {
	var requestBody dto.NewUserRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		errBindJSON := errs.NewUnprocessibleEntityError("invalid request body")
		ctx.JSON(errBindJSON.Status(), errBindJSON)
		return
	}

	result, err := uh.userService.CreateNewUser(requestBody)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

func (uh *userHandler) Login(ctx *gin.Context) {
	var requestBody dto.LoginUserRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		errBindJSON := errs.NewUnprocessibleEntityError("invalid request body")
		ctx.JSON(errBindJSON.Status(), errBindJSON)
		return
	}

	result, err := uh.userService.Login(requestBody)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (uh *userHandler) UpdateUser(ctx *gin.Context) {
	var requestBody dto.UpdateUserRequest

	userData, ok := ctx.MustGet("userData").(*entity.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.Status(), newError)
		return
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		errBindJSON := errs.NewUnprocessibleEntityError("invalid request body")
		ctx.JSON(errBindJSON.Status(), errBindJSON)
		return
	}

	result, err := uh.userService.UpdateUser(userData, &requestBody)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (uh *userHandler) DeleteUser(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(*entity.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.Status(), newError)
		return
	}

	result, err := uh.userService.DeleteUser(userData)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
