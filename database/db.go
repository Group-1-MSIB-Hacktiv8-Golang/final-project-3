package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "Indhiro123"
	dbname   = "kanban_board"
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

func handleCreateRequiredTables() {
	userTable := `
		CREATE TABLE IF NOT EXISTS "user" (
			id SERIAL PRIMARY KEY,
			full_name TEXT NOT NULL,
			email TEXT NOT NULL,
			password TEXT NOT NULL,
			role TEXT NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now()
		);
	`

	_, err = db.Exec(userTable)

	if err != nil {
		log.Panic("error occured while trying to create required tables", err)
	}
}

func InitiliazeDatabase() {
	handleDatabaseConnection()
	handleCreateRequiredTables()
}

func GetDatabaseInstance() *sql.DB {
	return db
}
