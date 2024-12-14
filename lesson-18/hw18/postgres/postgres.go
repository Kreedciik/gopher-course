package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "goonline"
	password = "root"
)

func InitDB() (*sql.DB, error) {
	dbConnection := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password,
	)
	db, err := sql.Open("postgres", dbConnection)
	if err != nil {
		return nil, err
	}
	return db, nil
}
