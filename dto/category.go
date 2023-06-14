package dto

import (
	"agolang/project-3/entity"
	"time"
)

/////////////////////////////////

type NewCategoryRequest struct {
	Type string `json:"type" valid:"required~type cannot be empty"`
}

func (c *NewCategoryRequest) ToEntity() *entity.Category {
	return &entity.Category{
		Type: c.Type,
	}
}

type NewCategoryResponse struct {
	Id        int       `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

/////////////////////////////////

type GetAllCategoriesResponse struct {
	Id        int        `json:"id"`
	Type      string     `json:"type"`
	UpdatedAt time.Time  `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
	Tasks     []TaskData `json:"tasks"`
}

/////////////////////////////////

type UpdateCategoryRequest NewCategoryRequest

func (c *UpdateCategoryRequest) ToEntity() *entity.Category {
	return &entity.Category{
		Type: c.Type,
	}
}

type UpdateCategoryResponse struct {
	Id        int       `json:"id"`
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
}

/////////////////////////////////

type DeleteCategoryResponse struct {
	Message string `json:"message"`
}

/////////////////////////////////
