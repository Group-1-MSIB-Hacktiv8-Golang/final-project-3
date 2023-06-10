package user_pg

import (
	"agolang/project-3/entity"
	"agolang/project-3/pkg/errs"
	"agolang/project-3/repository/user_repository"
	"database/sql"
	"errors"
	"fmt"
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
		SELECT id, email, password from "user"
		WHERE id = $1;
	`
	updateUserQuery = `
		UPDATE "user"
		SET full_name = $2,
		email = $3
		WHERE id = $1
		RETURNING id, full_name, email, updated_at;
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
	role := "member"

	var user entity.User

	row := u.db.QueryRow(createNewUser, payload.FullName, payload.Email, payload.Password, role)

	err := row.Scan(&user.Id, &user.FullName, &user.Email, &user.CreatedAt)

	if err != nil {
		return nil, errs.NewInternalServerError(fmt.Errorf("failed to create new user: %w", err).Error())
	}

	return &user, nil
}

func (u *userPG) GetUserById(userId int) (*entity.User, errs.MessageErr) {
	var user entity.User

	row := u.db.QueryRow(retrieveUserById, userId)

	err := row.Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}

func (u *userPG) GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr) {
	var user entity.User

	tx, err := u.db.Begin()

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	row := tx.QueryRow(retrieveUserByEmail, userEmail)

	err = row.Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError(fmt.Errorf("failed to create new user: %w", err).Error())
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}

func (u *userPG) UpdateUser(user *entity.User, payload *entity.User) (*entity.User, errs.MessageErr) {

	tx, err := u.db.Begin()

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	row := tx.QueryRow(updateUserQuery, payload.Id, payload.FullName, payload.Email)

	err = row.Scan(&user.Id, &user.FullName, &user.Email, &user.Password)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError(fmt.Errorf("user not found: %w", err).Error())
		}
		return nil, errs.NewInternalServerError(fmt.Errorf("failed to update user: %w", err).Error())
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return user, nil
}
