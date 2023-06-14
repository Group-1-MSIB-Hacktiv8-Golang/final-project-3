package user_repository

import (
	"agolang/project-3/entity"
	"agolang/project-3/pkg/errs"
)

type UserRepository interface {
	CreateNewUser(payload entity.User) (*entity.User, errs.MessageErr)
	GetUserById(userId int) (*entity.User, errs.MessageErr)
	GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr)
	UpdateUser(newUser *entity.User, userId int) (*entity.User, errs.MessageErr)
	DeleteUser(userId int) errs.MessageErr
}
