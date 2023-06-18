package task_repository

import (
	"agolang/project-3/entity"
	"agolang/project-3/pkg/errs"
)

type TaskRepository interface {
	CreateTask(user *entity.User, task *entity.Task) (*entity.Task, errs.MessageErr)
	GetAllTasks() ([]entity.Task, errs.MessageErr)
	GetAllTasksByCategoryId(categoryId int) ([]entity.Task, errs.MessageErr)
	GetTaskById(id int) (*entity.Task, errs.MessageErr)
	UpdateTask(oldTask *entity.Task, newTask *entity.Task) (*entity.Task, errs.MessageErr)
	UpdateTaskStatus(oldTask *entity.Task, newTask *entity.Task) (*entity.Task, errs.MessageErr)
	UpdateTaskCategory(id int, newCategoryId int) (*entity.Task, errs.MessageErr)
	DeleteTask(id int) errs.MessageErr
}
