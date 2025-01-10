package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Postgres struct {
	Db *sql.DB
}

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
	SSLMode  string
}

func NewPostgres(cfg Config) (*sql.DB, error) {
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User,
		cfg.Password, cfg.Dbname,
		cfg.SSLMode)

	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
