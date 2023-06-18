package task_pg

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"agolang/project-3/entity"
	"agolang/project-3/pkg/errs"
	"agolang/project-3/repository/task_repository"
)

const (
	createNewTask = `
	INSERT INTO "task"
	(
		title,
		description,
		status,
		user_id,
		category_id
	)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, title, status, description, user_id, category_id, created_at
	`
	getAllTasks = `
		SELECT id, title, status, description, user_id, category_id, created_at
		FROM "task"
	`
	getTaskById = `
		SELECT id, title, status, description, user_id, category_id, created_at, updated_at
		FROM "task"
		WHERE id = $1
	`
	updateTask = `
		UPDATE "task"
		SET title = $2,
		description = $3,
		updated_at = $4
		WHERE id = $1
		RETURNING id, title, description, status, user_id, category_id, updated_at
	`
	UpdateTaskStatus = `
		UPDATE "task"
		SET status = $2,
		updated_at = $3
		WHERE id = $1
		RETURNING id, title, description, status, user_id, category_id, updated_at
	`
	UpdateTaskCategory = `
		UPDATE "task"
		SET category_id = $2,
		updated_at = $3
		WHERE id = $1
		RETURNING id, title, description, status, user_id, category_id, updated_at
	`
	DeleteTask = `
		DELETE FROM "task"
		WHERE id = $1;
	`
)

type taskPG struct {
	db *sql.DB
}

func NewTaskPG(db *sql.DB) task_repository.TaskRepository {
	return &taskPG{
		db: db,
	}
}

func (t *taskPG) CreateTask(user *entity.User, task *entity.Task) (*entity.Task, errs.MessageErr) {

	task.Status = false

	var tasks entity.Task

	row := t.db.QueryRow(createNewTask, task.Title, task.Description, task.Status, user.Id, task.CategoryId)

	err := row.Scan(&tasks.Id, &tasks.Title, &tasks.Status, &tasks.Description, &tasks.UserId, &tasks.CategoryId, &tasks.CreatedAt)

	if err != nil {
		return nil, errs.NewNotFoundError("Failed to create new task")
	}

	return &tasks, nil
}

func (t *taskPG) GetAllTasks() ([]entity.Task, errs.MessageErr) {
	var tasks []entity.Task

	rows, err := t.db.Query(getAllTasks)

	if err != nil {
		return nil, errs.NewInternalServerError("Failed to get all task")
	}

	defer rows.Close()

	for rows.Next() {
		task := entity.Task{}
		err = rows.Scan(&task.Id, &task.Title, &task.Status, &task.Description, &task.UserId, &task.CategoryId, &task.CreatedAt)

		if err != nil {
			return nil, errs.NewInternalServerError("Failed to get all task")
		}
		// return nil, errs.NewInternalServerError(fmt.Errorf("failed to create new task: %w", err).Error())

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (t *taskPG) GetAllTasksByCategoryId(categoryId int) ([]entity.Task, errs.MessageErr) {
	var tasks []entity.Task

	return tasks, nil
}

func (t *taskPG) GetTaskById(id int) (*entity.Task, errs.MessageErr) {
	var task entity.Task

	row := t.db.QueryRow(getTaskById, id)

	err := row.Scan(&task.Id, &task.Title, &task.Status, &task.Description, &task.UserId, &task.CategoryId, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("Task not found")
		}
		return nil, errs.NewInternalServerError("Something went wrong")
	}

	return &task, nil
}

func (t *taskPG) UpdateTask(oldTask *entity.Task, newTask *entity.Task) (*entity.Task, errs.MessageErr) {

	newTask.UpdatedAt = time.Now()

	row := t.db.QueryRow(updateTask, oldTask.Id, newTask.Title, newTask.Description, newTask.UpdatedAt)

	err := row.Scan(&oldTask.Id, &oldTask.Title, &oldTask.Description, &oldTask.Status, &oldTask.UserId, &oldTask.CategoryId, &oldTask.UpdatedAt)

	if err != nil {
		return nil, errs.NewNotFoundError("Failed to update task")
	}

	return oldTask, nil
}

func (t *taskPG) UpdateTaskStatus(oldTask *entity.Task, newTask *entity.Task) (*entity.Task, errs.MessageErr) {

	oldTask.UpdatedAt = time.Now()

	oldTask.Status = newTask.Status

	row := t.db.QueryRow(UpdateTaskStatus, oldTask.Id, oldTask.Status, oldTask.UpdatedAt)

	err := row.Scan(&oldTask.Id, &oldTask.Title, &oldTask.Description, &oldTask.Status, &oldTask.UserId, &oldTask.CategoryId, &oldTask.UpdatedAt)

	if err != nil {
		return nil, errs.NewNotFoundError(fmt.Errorf("failed to create new task: %w", err).Error())
	}

	return oldTask, nil
}

func (t *taskPG) UpdateTaskCategory(id int, newCategoryId int) (*entity.Task, errs.MessageErr) {
	task, err2 := t.GetTaskById(id)
	if err2 != nil {
		return nil, err2
	}

	task.CategoryId = newCategoryId

	task.UpdatedAt = time.Now()

	row := t.db.QueryRow(UpdateTaskCategory, task.Id, task.CategoryId, task.UpdatedAt)

	fmt.Printf("%+v", task)

	err := row.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.UserId, &task.CategoryId, &task.UpdatedAt)
	fmt.Println("err di update task", err)

	if err != nil {
		return nil, errs.NewNotFoundError("Failed to update task")
	}

	return task, nil
}

func (t *taskPG) DeleteTask(id int) errs.MessageErr {
	_, err := t.db.Exec(DeleteTask, id)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return errs.NewNotFoundError("User not found")
		}
		return errs.NewInternalServerError("Failed to delete user")
	}

	return nil
}
