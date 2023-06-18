package service

import (
	"agolang/project-3/dto"
	"agolang/project-3/entity"
	"agolang/project-3/pkg/errs"
	"agolang/project-3/repository/category_repository"
	"agolang/project-3/repository/task_repository"
	"agolang/project-3/repository/user_repository"
)

type TaskService interface {
	CreateTask(user *entity.User, payload *dto.NewTaskRequest) (*dto.NewTaskResponse, errs.MessageErr)
	GetAllTasks() ([]dto.GetAllTasksResponse, errs.MessageErr)
	UpdateTask(id int, payload *dto.UpdateTaskRequest) (*dto.UpdateTaskResponse, errs.MessageErr)
	UpdateTaskStatus(id int, payload *dto.UpdateTaskStatusRequest) (*dto.UpdateTaskResponse, errs.MessageErr)
	UpdateTaskCategory(id int, payload *dto.UpdateTaskCategoryRequest) (*dto.UpdateTaskResponse, errs.MessageErr)
	DeleteTask(id int) (*dto.DeleteTaskResponse, errs.MessageErr)
}

type taskService struct {
	taskRepo     task_repository.TaskRepository
	categoryRepo category_repository.CategoryRepository
	userRepo     user_repository.UserRepository
}

func NewTaskService(taskRepo task_repository.TaskRepository, categoryRepo category_repository.CategoryRepository, userRepo user_repository.UserRepository) TaskService {
	return &taskService{taskRepo, categoryRepo, userRepo}
}

func (t *taskService) CreateTask(user *entity.User, payload *dto.NewTaskRequest) (*dto.NewTaskResponse, errs.MessageErr) {
	task := payload.ToEntity()

	_, err := t.categoryRepo.GetCategoryById(task.CategoryId)
	if err != nil {
		return nil, errs.NewNotFoundError("Failed to create task")
	}

	createdTask, err := t.taskRepo.CreateTask(user, task)
	if err != nil {
		return nil, errs.NewNotFoundError("Failed to create task")
	}

	response := &dto.NewTaskResponse{
		Id:          createdTask.Id,
		Title:       createdTask.Title,
		Status:      createdTask.Status,
		Description: createdTask.Description,
		UserId:      createdTask.UserId,
		CategoryId:  createdTask.CategoryId,
		CreatedAt:   createdTask.CreatedAt,
	}

	return response, nil
}

func (t *taskService) GetAllTasks() ([]dto.GetAllTasksResponse, errs.MessageErr) {
	tasks, err := t.taskRepo.GetAllTasks()
	if err != nil {
		return nil, errs.NewNotFoundError("Failed to get all task")
	}

	response := []dto.GetAllTasksResponse{}
	for _, task := range tasks {
		user, err := t.userRepo.GetUserById(task.UserId)
		if err != nil {
			return nil, errs.NewNotFoundError("Failed to get all task")
		}
		response = append(response, dto.GetAllTasksResponse{
			Id:          task.Id,
			Title:       task.Title,
			Status:      task.Status,
			Description: task.Description,
			UserId:      task.UserId,
			CategoryId:  task.CategoryId,
			CreatedAt:   task.CreatedAt,
			User: dto.UserData{
				Id:       user.Id,
				Email:    user.Email,
				FullName: user.FullName,
			},
		})
	}

	return response, nil
}

func (t *taskService) UpdateTask(id int, payload *dto.UpdateTaskRequest) (*dto.UpdateTaskResponse, errs.MessageErr) {

	oldTask, err := t.taskRepo.GetTaskById(id)
	if err != nil {
		return nil, errs.NewNotFoundError("Failed to update task")
	}

	newTask := payload.ToEntity()

	updatedTask, err2 := t.taskRepo.UpdateTask(oldTask, newTask)
	if err2 != nil {
		return nil, err2
	}

	response := &dto.UpdateTaskResponse{
		Id:          updatedTask.Id,
		Title:       updatedTask.Title,
		Description: updatedTask.Description,
		Status:      updatedTask.Status,
		UserId:      updatedTask.UserId,
		CategoryId:  updatedTask.CategoryId,
		UpdatedAt:   updatedTask.UpdatedAt,
	}

	return response, nil
}

func (t *taskService) UpdateTaskStatus(id int, payload *dto.UpdateTaskStatusRequest) (*dto.UpdateTaskResponse, errs.MessageErr) {
	oldTask, err := t.taskRepo.GetTaskById(id)
	if err != nil {
		return nil, errs.NewNotFoundError("Failed to update status task")
	}

	newTask := payload.ToEntity()

	updatedTask, err2 := t.taskRepo.UpdateTaskStatus(oldTask, newTask)
	if err2 != nil {
		return nil, errs.NewNotFoundError("Failed to update status task")
	}

	response := &dto.UpdateTaskResponse{
		Id:          updatedTask.Id,
		Title:       updatedTask.Title,
		Description: updatedTask.Description,
		Status:      updatedTask.Status,
		UserId:      updatedTask.UserId,
		CategoryId:  updatedTask.CategoryId,
		UpdatedAt:   updatedTask.UpdatedAt,
	}

	return response, nil
}

func (t *taskService) UpdateTaskCategory(id int, payload *dto.UpdateTaskCategoryRequest) (*dto.UpdateTaskResponse, errs.MessageErr) {
	_, err := t.categoryRepo.GetCategoryById(payload.CategoryId)

	if err != nil {
		return nil, errs.NewNotFoundError("Failed to update task category")
	}

	updatedCategory, err := t.taskRepo.UpdateTaskCategory(id, payload.CategoryId)
	if err != nil {
		return nil, errs.NewNotFoundError("Failed to update task category")
	}

	response := &dto.UpdateTaskResponse{
		Id:          updatedCategory.Id,
		Title:       updatedCategory.Title,
		Description: updatedCategory.Description,
		Status:      updatedCategory.Status,
		UserId:      updatedCategory.UserId,
		CategoryId:  updatedCategory.CategoryId,
		UpdatedAt:   updatedCategory.UpdatedAt,
	}

	return response, nil
}

func (t *taskService) DeleteTask(id int) (*dto.DeleteTaskResponse, errs.MessageErr) {
	err := t.taskRepo.DeleteTask(id)
	if err != nil {
		return nil, errs.NewNotFoundError("Failed to delete task")
	}

	response := &dto.DeleteTaskResponse{
		Message: "Task has been successfully deleted",
	}

	return response, nil
}
