package category_repository

import (
	"agolang/project-3/entity"
	"agolang/project-3/pkg/errs"
)

type CategoryRepository interface {
	CreateCategory(payload entity.Category) (*entity.Category, errs.MessageErr)
	GetAllCategories() ([]CategoryTask, errs.MessageErr)
	GetCategoryById(id int) ([]entity.Category, errs.MessageErr)
	UpdateCategory(oldCategory *entity.Category, userId int) (*entity.Category, errs.MessageErr)
	DeleteCategory(categoryId int) errs.MessageErr
}
