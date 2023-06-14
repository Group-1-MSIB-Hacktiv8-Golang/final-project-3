package dto

import (
	"time"

	"agolang/project-3/entity"
)

/////////////////////////////////

type TaskData struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserId      int       `json:"user_id"`
	CategoryId  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

/////////////////////////////////

type NewTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	CategoryId  int    `json:"category_id" `
}

func (t *NewTaskRequest) ToEntity() *entity.Task {
	return &entity.Task{
		Title:       t.Title,
		Description: t.Description,
		CategoryId:  t.CategoryId,
	}
}

type NewTaskResponse struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Status      bool      `json:"status"`
	Description string    `json:"description"`
	UserId      int       `json:"user_id"`
	CategoryId  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}

/////////////////////////////////

type GetAllTasksResponse struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Status      bool      `json:"status"`
	Description string    `json:"description"`
	UserId      int       `json:"user_id"`
	CategoryId  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	User        UserData  `json:"user"`
}

/////////////////////////////////

type UpdateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (t *UpdateTaskRequest) ToEntity() *entity.Task {
	return &entity.Task{
		Title:       t.Title,
		Description: t.Description,
	}
}

type UpdateTaskResponse struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserId      int       `json:"user_id"`
	CategoryId  int       `json:"category_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}

/////////////////////////////////

type UpdateTaskStatusRequest struct {
	Status bool `json:"status"`
}

/////////////////////////////////

type UpdateTaskCategoryRequest struct {
	CategoryId int `json:"category_id"`
}

/////////////////////////////////

type DeleteTaskResponse struct {
	Message string `json:"message"`
}
