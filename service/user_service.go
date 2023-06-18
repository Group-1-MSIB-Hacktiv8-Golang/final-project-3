package service

import (
	"agolang/project-3/dto"
	"agolang/project-3/entity"
	"agolang/project-3/pkg/errs"
	"agolang/project-3/pkg/helpers"
	"agolang/project-3/repository/user_repository"
	"net/http"
)

type UserService interface {
	CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr)
	Login(loginUserRequest dto.LoginUserRequest) (*dto.LoginUserResponse, errs.MessageErr)
	UpdateUser(user *entity.User, payload *dto.UpdateUserRequest) (*dto.UpdateUserResponse, errs.MessageErr)
	DeleteUser(user *entity.User) (*dto.DeleteUserResponse, errs.MessageErr)
}

type userService struct {
	userRepo user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	if len(payload.Password) < 6 {
		return nil, errs.NewInternalServerError("Password cant be less then 6")
	}

	user := entity.User{
		FullName: payload.FullName,
		Email:    payload.Email,
		Password: payload.Password,
	}

	err = user.HashPassword()

	if err != nil {
		return nil, err
	}

	var entityUser *entity.User

	//Memasukan data user kedalam table
	entityUser, err = u.userRepo.CreateNewUser(user)

	if err != nil {
		return nil, err
	}

	response := dto.NewUserResponse{
		Id:        entityUser.Id,
		FullName:  entityUser.FullName,
		Email:     entityUser.Email,
		CreatedAt: entityUser.CreatedAt,
	}

	return &response, nil
}

func (u *userService) Login(loginUserRequest dto.LoginUserRequest) (*dto.LoginUserResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(loginUserRequest)

	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.GetUserByEmail(loginUserRequest.Email)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewBadRequest("invalid email/password")
		}
		return nil, err
	}

	isValidPassword := user.ComparePassword(loginUserRequest.Password)

	if !isValidPassword {
		return nil, errs.NewBadRequest("invalid email/password")
	}

	token := user.GenerateToken()

	response := dto.LoginUserResponse{
		Token: token,
	}

	return &response, nil
}

func (u *userService) UpdateUser(user *entity.User, payload *dto.UpdateUserRequest) (*dto.UpdateUserResponse, errs.MessageErr) {

	newUser := payload.ToEntity()

	updatedUser, err := u.userRepo.UpdateUser(newUser, user.Id)
	if err != nil {
		return nil, err
	}

	response := &dto.UpdateUserResponse{
		Id:        updatedUser.Id,
		FullName:  updatedUser.FullName,
		Email:     updatedUser.Email,
		UpdatedAt: updatedUser.UpdatedAt,
	}

	return response, nil
}

func (u *userService) DeleteUser(user *entity.User) (*dto.DeleteUserResponse, errs.MessageErr) {
	if err := u.userRepo.DeleteUser(user.Id); err != nil {
		return nil, err
	}

	response := &dto.DeleteUserResponse{
		Message: "Your account has been succesfully deleted",
	}

	return response, nil
}
