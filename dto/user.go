package dto

import (
	"agolang/project-3/entity"
	"time"
)

/////////////////////////////////

type NewUserRequest struct {
	FullName string `json:"full_name" valid:"required~full_name cannot be empty"`
	Email    string `json:"email" valid:"required~email cannot be empty"`
	Password string `json:"password" valid:"required~password cannot be empty"`
}

type NewUserResponse struct {
	StatusCode int       `json:"statusCode"`
	Id         int       `json:"id"`
	FullName   string    `json:"full_name"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"created_at"`
}

/////////////////////////////////

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type LoginUserResponse struct {
	StatusCode int           `json:"statusCode"`
	Data       TokenResponse `json:"data"`
}

/////////////////////////////////

type UpdateUserRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

func (r *UpdateUserRequest) ToEntity() *entity.User {
	return &entity.User{
		FullName: r.FullName,
		Email:    r.Email,
	}
}

type UpdateUserResponse struct {
	StatusCode int       `json:"statusCode"`
	Id         int       `json:"id"`
	FullName   string    `json:"full_name"`
	Email      string    `json:"email"`
	UpdatedAt  time.Time `json:"updated_at"`
}

/////////////////////////////////

type DeleteUserResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

/////////////////////////////////

// ?
type UserData struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}
