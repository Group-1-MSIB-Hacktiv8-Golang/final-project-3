package user_pg

import (
	"agolang/project-3/entity"
	"agolang/project-3/pkg/errs"
	"agolang/project-3/repository/user_repository"
	"database/sql"
	"errors"
	"time"
)

const (
	createNewUser = `
		INSERT INTO "user"
		(
			full_name,
			email,
			password,
			role
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id, full_name, email, created_at
	`
	//Dont forget "" at table name
	retrieveUserByEmail = `
		SELECT id, email, password from "user"
		WHERE email = $1;
	`
	retrieveUserById = `
		SELECT id, full_name, email, password, role, created_at, updated_at from "user"
		WHERE id = $1;
	`
	updateUserQuery = `
		UPDATE "user"
		SET full_name = $2,
		email = $3,
		updated_at = $4
		WHERE id = $1
		RETURNING id, full_name, email, updated_at
	`
	deleteUserQuery = `
		DELETE FROM "user"
		WHERE id = $1;
	`
)

type userPG struct {
	db *sql.DB
}

func NewUserPG(db *sql.DB) user_repository.UserRepository {
	return &userPG{
		db: db,
	}
}

func (u *userPG) CreateNewUser(payload entity.User) (*entity.User, errs.MessageErr) {
	payload.Role = "member"

	var user entity.User

	row := u.db.QueryRow(createNewUser, payload.FullName, payload.Email, payload.Password, payload.Role)

	err := row.Scan(&user.Id, &user.FullName, &user.Email, &user.CreatedAt)

	if err != nil {
		return nil, errs.NewInternalServerError("Email has been used")
	}

	return &user, nil
}

func (u *userPG) GetUserById(userId int) (*entity.User, errs.MessageErr) {
	var user entity.User

	row := u.db.QueryRow(retrieveUserById, userId)

	err := row.Scan(&user.Id, &user.FullName, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("User not found")
		}
		return nil, errs.NewInternalServerError("Something went wrong")
	}

	return &user, nil
}

func (u *userPG) GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr) {
	var user entity.User

	tx, err := u.db.Begin()

	if err != nil {
		return nil, errs.NewInternalServerError("Something went wrong")
	}

	row := tx.QueryRow(retrieveUserByEmail, userEmail)

	err = row.Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("User not found")
		}
		return nil, errs.NewInternalServerError("Failed to create new user")
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("Something went wrong")
	}

	return &user, nil
}

func (u *userPG) UpdateUser(payload *entity.User, userId int) (*entity.User, errs.MessageErr) {

	var user entity.User

	payload.UpdatedAt = time.Now()

	tx, err := u.db.Begin()

	if err != nil {
		return nil, errs.NewInternalServerError("Something went wrong")
	}

	row := tx.QueryRow(updateUserQuery, userId, payload.FullName, payload.Email, payload.UpdatedAt)

	err = row.Scan(&user.Id, &user.FullName, &user.Email, &user.UpdatedAt)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("User not found")
		}
		return nil, errs.NewInternalServerError("Failed to update user")
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("Something went wrong")
	}

	return &user, nil
}

func (u *userPG) DeleteUser(userId int) errs.MessageErr {

	_, err := u.db.Exec(deleteUserQuery, userId)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return errs.NewNotFoundError("User not found")
		}
		return errs.NewInternalServerError("Failed to delete user")
	}

	return nil
}
