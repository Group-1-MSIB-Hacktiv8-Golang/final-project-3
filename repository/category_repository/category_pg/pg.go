package category_pg

import (
	"agolang/project-3/entity"
	"agolang/project-3/pkg/errs"
	"agolang/project-3/repository/category_repository"
	"database/sql"
	"errors"
	"time"
)

const (
	createNewCategory = `
		INSERT INTO "category"
		(
			type
		)
		VALUES ($1)
		RETURNING id, type, created_at
	`
	getAllCategory = `
		SELECT id, type, updated_at, created_at from "category"
	`
	getCategoryByIdQuery = `
		SELECT id, type, updated_at, created_at from "category"
		WHERE id = $1;
	`
	updateCategory = `
		UPDATE "category"
		SET type = $2,
		updated_at = $3
		WHERE id = $1
		RETURNING id, type, updated_at
	`
	deleteCategoryQuery = `
		DELETE FROM "category"
		WHERE id = $1;
	`
	getTaskByCategoryId = `
		SELECT id, title, description, user_id, category_id, created_at, updated_at
		FROM "task"
		WHERE category_id = $1
	`
)

type categoryPG struct {
	db *sql.DB
}

func NewCategoryPG(db *sql.DB) category_repository.CategoryRepository {
	return &categoryPG{
		db: db,
	}
}

func (c *categoryPG) getTaskByCategoryId(category_id int) ([]entity.Task, errs.MessageErr) {
	rows, err := c.db.Query(getTaskByCategoryId, category_id)
	if err != nil {
		return nil, errs.NewInternalServerError("failed to get task")
	}
	defer rows.Close()

	tasks := []entity.Task{}

	for rows.Next() {
		task := entity.Task{}

		err = rows.Scan(&task.Id, &task.Title, &task.Description, &task.UserId, &task.CategoryId, &task.CreatedAt, &task.UpdatedAt)

		if err != nil {
			return nil, errs.NewInternalServerError("failed to get task")
		}

		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (c *categoryPG) CreateCategory(payload entity.Category) (*entity.Category, errs.MessageErr) {
	var category entity.Category

	row := c.db.QueryRow(createNewCategory, payload.Type)

	err := row.Scan(&category.Id, &category.Type, &category.CreatedAt)

	if err != nil {
		return nil, errs.NewInternalServerError("failed to create new category")
	}

	return &category, nil
}

func (c *categoryPG) GetAllCategories() ([]category_repository.CategoryTask, errs.MessageErr) {
	var category entity.Category

	var categories []category_repository.CategoryTask

	rows, err := c.db.Query(getAllCategory)

	if err != nil {
		return nil, errs.NewInternalServerError("failed to get all category")
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&category.Id, &category.Type, &category.CreatedAt, &category.UpdatedAt)

		if err != nil {
			return nil, errs.NewInternalServerError("failed to get all category")
		}

		tasks, err := c.getTaskByCategoryId(category.Id)

		if err != nil {
			return nil, errs.NewInternalServerError("failed to get all category")
		}

		categoryTask := category_repository.CategoryTask{
			Category: category,
			Tasks:    tasks,
		}

		categories = append(categories, categoryTask)
	}

	return categories, nil
}

func (c *categoryPG) GetCategoryById(id int) ([]entity.Category, errs.MessageErr) {
	var category entity.Category

	categories := []entity.Category{}

	rows, err := c.db.Query(getCategoryByIdQuery, id)

	if err != nil {
		return nil, errs.NewInternalServerError("failed to create new category")
	}

	for rows.Next() {
		err = rows.Scan(&category.Id, &category.Type, &category.UpdatedAt, &category.CreatedAt)

		if err != nil {
			return nil, errs.NewInternalServerError("failed to create new category")
		}

		categories = append(categories, category)
	}

	if err != nil {
		return nil, errs.NewInternalServerError("failed to create new category")
	}

	return categories, nil

}

func (c *categoryPG) UpdateCategory(oldCategory *entity.Category, categoryId int) (*entity.Category, errs.MessageErr) {
	var category entity.Category

	oldCategory.UpdatedAt = time.Now()

	tx, err := c.db.Begin()

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	row := tx.QueryRow(updateCategory, categoryId, oldCategory.Type, oldCategory.UpdatedAt)

	err = row.Scan(&category.Id, &category.Type, &category.UpdatedAt)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("category not found")
		}
		return nil, errs.NewInternalServerError("failed to update category")
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return oldCategory, nil

}
func (c *categoryPG) DeleteCategory(categoryId int) errs.MessageErr {

	_, err := c.db.Exec(deleteCategoryQuery, categoryId)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return errs.NewNotFoundError("Category not found")
		}
		return errs.NewInternalServerError("Failed to update category")
	}

	return nil

}
