package database

import (
	"agolang/project-3/entity"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "containers-us-west-122.railway.app"
	port     = "7926"
	user     = "postgres"
	password = "bZA60fYIF0U0WCTxTYw7"
	dbname   = "railway"
	dialect  = "postgres"
)

var (
	db  *sql.DB
	err error
)

func handleDatabaseConnection() {
	//Dialect tidak di input
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err = sql.Open(dialect, psqlInfo)

	if err != nil {
		log.Panic("error occured while trying to validate database arguments:", err)
	}
}

// type UpdateUserRequest struct {
// 	Email    string `json:"email" valid:"required~email cannot be empty,email~email must be a valid email"`
// 	Username string `json:"username" valid:"required~username cannot be empty"`
// }

func handleCreateRequiredTables() {
	userTable := `
		CREATE TABLE IF NOT EXISTS "user" (
			id SERIAL PRIMARY KEY,
			full_name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			role TEXT NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now()
		);
	`
	categoryTable := `
		CREATE TABLE IF NOT EXISTS "category" (
			id SERIAL PRIMARY KEY,
			type TEXT NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now()
		);
	`
	taskTable := `
		CREATE TABLE IF NOT EXISTS "task" (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			description TEXT NOT NULL,
			status TEXT NOT NULL,
			user_id int NOT NULL,
			category_id int NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now(),
			CONSTRAINT task_user_id_fk
			FOREIGN KEY(user_id) REFERENCES "user"(id) ON DELETE CASCADE,
			CONSTRAINT task_category_id_fk
			FOREIGN KEY(category_id) REFERENCES "category"(id) ON DELETE CASCADE	
		);
	`
	seedAdmin := `
		INSERT INTO "user"
		(
			full_name,
			email,
			password,
			role
		)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (email) do nothing
	`
	_, err = db.Exec(userTable)

	if err != nil {
		log.Panic("error occured while trying to create user table", err)
	}

	_, err = db.Exec(categoryTable)

	if err != nil {
		log.Panic("error occured while trying to create category table", err)
	}

	_, err = db.Exec(taskTable)

	if err != nil {
		log.Panic("error occured while trying to create task table", err)
	}

	user := entity.User{
		FullName: "admin",
		Email:    "admin@gmail.com",
		Password: "admin123",
		Role:     "admin",
	}

	err = user.HashPassword()

	_, err = db.Exec(seedAdmin, user.FullName, user.Email, user.Password, user.Role)

	if err != nil {
		log.Panic("error occured while trying to seeding admin data", err)
	}
}

func InitiliazeDatabase() {
	handleDatabaseConnection()
	handleCreateRequiredTables()
}

func GetDatabaseInstance() *sql.DB {
	return db
}
