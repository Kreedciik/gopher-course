package repository

import "database/sql"

type Repository struct {
	User *UserRepository
}

func NewRepository(postgresDB *sql.DB) *Repository {
	return &Repository{
		User: NewUserRepository(postgresDB),
	}
}
