package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"agolang/project-3/dto"
	"agolang/project-3/entity"
	"agolang/project-3/pkg/errs"
	"agolang/project-3/service"
)

type TaskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) *TaskHandler {
	return &TaskHandler{taskService}
}

func (t *TaskHandler) CreateTask(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(*entity.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.Status(), newError)
		return
	}
	var requestBody dto.NewTaskRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessibleEntityError(err.Error())
		ctx.JSON(newError.Status(), newError)
		return
	}

	createdTask, err := t.taskService.CreateTask(userData, &requestBody)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusCreated, createdTask)
}

func (t *TaskHandler) GetAllTasks(ctx *gin.Context) {
	tasks, err := t.taskService.GetAllTasks()
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (t *TaskHandler) UpdateTask(ctx *gin.Context) {
	taskId := ctx.Param("taskId")
	taskIdInt, err := strconv.ParseUint(taskId, 10, 32)
	if err != nil {
		errValidation := errs.NewBadRequest("Task id should be in unsigned integer")
		ctx.JSON(errValidation.Status(), errValidation)
		return
	}

	var reqBody dto.UpdateTaskRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		errValidation := errs.NewUnprocessibleEntityError(err.Error())
		ctx.JSON(errValidation.Status(), errValidation)
		return
	}

	updatedTask, errUpdate := t.taskService.UpdateTask(int(taskIdInt), &reqBody)
	if errUpdate != nil {
		ctx.JSON(errUpdate.Status(), errUpdate)
		return
	}

	ctx.JSON(http.StatusOK, updatedTask)
}

func (t *TaskHandler) UpdateTaskStatus(ctx *gin.Context) {
	taskId := ctx.Param("taskId")
	taskIdInt, err := strconv.ParseUint(taskId, 10, 32)
	if err != nil {
		errValidation := errs.NewBadRequest("Task id should be in unsigned integer")
		ctx.JSON(errValidation.Status(), errValidation)
		return
	}

	var reqBody dto.UpdateTaskStatusRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		validationError := errs.NewUnprocessibleEntityError(err.Error())
		ctx.JSON(validationError.Status(), validationError)
		return
	}

	fmt.Println(int(taskIdInt))

	response, err := t.taskService.UpdateTaskStatus(int(taskIdInt), &reqBody)

	fmt.Println(response)

	ctx.JSON(http.StatusOK, response)
}

func (t *TaskHandler) UpdateTaskCategory(ctx *gin.Context) {
	taskId := ctx.Param("taskId")
	taskIdInt, err := strconv.ParseUint(taskId, 10, 32)
	if err != nil {
		errValidation := errs.NewBadRequest("Task id should be in unsigned integer")
		ctx.JSON(errValidation.Status(), errValidation)
		return
	}

	var reqBody dto.UpdateTaskCategoryRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		validationError := errs.NewUnprocessibleEntityError(err.Error())
		ctx.JSON(validationError.Status(), validationError)
		return
	}

	updatedCategory, errUpdate := t.taskService.UpdateTaskCategory(int(taskIdInt), &reqBody)
	if errUpdate != nil {
		ctx.JSON(errUpdate.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, updatedCategory)
}

func (t *TaskHandler) DeleteTask(ctx *gin.Context) {
	taskId := ctx.Param("taskId")
	taskIdInt, err := strconv.ParseUint(taskId, 10, 32)
	if err != nil {
		newError := errs.NewBadRequest("Task id should be in unsigned integer")
		ctx.JSON(newError.Status(), newError)
		return
	}

	response, err2 := t.taskService.DeleteTask(int(taskIdInt))
	if err2 != nil {
		ctx.JSON(err2.Status(), err2)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
