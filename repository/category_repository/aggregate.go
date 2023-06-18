package category_repository

import "agolang/project-3/entity"

type CategoryTask struct {
	Category entity.Category
	Tasks    []entity.Task
}
