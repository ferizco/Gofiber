package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "go_nico"
	password = "12345"
	dbname   = "golangfiber"
)

// ConnectDB connects to the PostgreSQL database and returns a *sql.DB instance
func ConnectDB() (*sql.DB, error) {

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
