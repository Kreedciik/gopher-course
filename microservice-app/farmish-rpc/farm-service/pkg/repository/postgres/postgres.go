package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
	SSLMode  string
}

func NewPostgres(cfg PostgresConfig) (*sql.DB, error) {
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User,
		cfg.Password, cfg.Dbname,
		cfg.SSLMode)
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return db, err
	}
	err = db.Ping()
	return db, err
}
