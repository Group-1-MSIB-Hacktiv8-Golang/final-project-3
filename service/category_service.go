package service

import (
	"agolang/project-3/dto"
	"agolang/project-3/entity"
	"agolang/project-3/pkg/errs"
	"agolang/project-3/pkg/helpers"
	"agolang/project-3/repository/category_repository"
	"agolang/project-3/repository/task_repository"
)

type CategoryService interface {
	CreateNewCategory(payload dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.MessageErr)
	GetAllCategories() ([]dto.GetAllCategoriesResponse, errs.MessageErr)
	UpdateCategory(categoryId int, payload *dto.UpdateCategoryRequest) (*dto.UpdateCategoryResponse, errs.MessageErr)
	DeleteCategory(categoryId int) (*dto.DeleteCategoryResponse, errs.MessageErr)
}

type categoryService struct {
	categoryRepo category_repository.CategoryRepository
	taskRepo     task_repository.TaskRepository
}

func NewCategoryService(categoryRepo category_repository.CategoryRepository, taskRepo task_repository.TaskRepository) CategoryService {
	return &categoryService{
		categoryRepo, taskRepo,
	}
}

func (c *categoryService) CreateNewCategory(payload dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	category := entity.Category{
		Type: payload.Type,
	}

	if err != nil {
		return nil, err
	}

	var entityCategory *entity.Category

	//Memasukan data category kedalam table
	entityCategory, err = c.categoryRepo.CreateCategory(category)

	if err != nil {
		return nil, err
	}

	response := dto.NewCategoryResponse{
		Id:        entityCategory.Id,
		Type:      entityCategory.Type,
		CreatedAt: entityCategory.CreatedAt,
	}

	return &response, nil
}

func (c *categoryService) GetAllCategories() ([]dto.GetAllCategoriesResponse, errs.MessageErr) {
	categories, err := c.categoryRepo.GetAllCategories()
	if err != nil {
		return nil, err
	}

	response := []dto.GetAllCategoriesResponse{}
	for _, category := range categories {
		tasksResponse := []dto.TaskData{}

		for _, task := range category.Tasks {

			tasksResponse = append(tasksResponse, dto.TaskData{
				Id:          task.Id,
				Title:       task.Title,
				Description: task.Description,
				UserId:      task.UserId,
				CategoryId:  task.CategoryId,
				CreatedAt:   task.CreatedAt,
				UpdatedAt:   task.UpdatedAt,
			})
		}

		categoryResponse := dto.GetAllCategoriesResponse{
			Id:        category.Category.Id,
			Type:      category.Category.Type,
			UpdatedAt: category.Category.UpdatedAt,
			CreatedAt: category.Category.CreatedAt,
			Tasks:     tasksResponse,
		}

		response = append(response, categoryResponse)

	}

	return response, nil
}

func (c *categoryService) UpdateCategory(categoryId int, payload *dto.UpdateCategoryRequest) (*dto.UpdateCategoryResponse, errs.MessageErr) {
	newCategory := payload.ToEntity()

	updateCategory, err := c.categoryRepo.UpdateCategory(newCategory, categoryId)

	if err != nil {
		return nil, err
	}

	response := &dto.UpdateCategoryResponse{
		Id:        categoryId,
		Type:      updateCategory.Type,
		UpdatedAt: updateCategory.UpdatedAt,
	}

	return response, nil
}

func (c *categoryService) DeleteCategory(categoryId int) (*dto.DeleteCategoryResponse, errs.MessageErr) {
	if err := c.categoryRepo.DeleteCategory(categoryId); err != nil {
		return nil, err
	}

	response := &dto.DeleteCategoryResponse{
		Message: "Category has been succesfully deleted",
	}

	return response, nil
}
