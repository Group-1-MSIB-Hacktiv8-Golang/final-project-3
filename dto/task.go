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
	Title       string `json:"title" valid:"required~title cannot be empty"`
	Description string `json:"description" valid:"required~description cannot be empty"`
	CategoryId  int    `json:"category_id" valid:"required~categoryId cannot be empty"`
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
	Title       string `json:"title" valid:"required~title cannot be empty"`
	Description string `json:"description" valid:"required~description cannot be empty"`
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
	Status bool `json:"status" valid:"required~status cannot be empty"`
}

func (t *UpdateTaskStatusRequest) ToEntity() *entity.Task {
	return &entity.Task{
		Status: t.Status,
	}
}

/////////////////////////////////

type UpdateTaskCategoryRequest struct {
	CategoryId int `json:"categoryId" valid:"required~categoryId cannot be empty"`
}

/////////////////////////////////

type DeleteTaskResponse struct {
	Message string `json:"message"`
}
