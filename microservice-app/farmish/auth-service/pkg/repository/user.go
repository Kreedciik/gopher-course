package repository

import (
	"database/sql"
	"farmish/model"

	"github.com/google/uuid"
)

type User interface {
	InsertUser(model.CreateUserDTO) error
	FindUserByEmail(string) (model.User, error)
}
type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) InsertUser(newUser model.CreateUserDTO) error {
	_, err := u.db.Exec(`
	INSERT INTO users VALUES ($1, $2, $3)
	`,
		uuid.NewString(),
		newUser.Email,
		newUser.Password,
	)
	return err
}

func (u *UserRepository) FindUserByEmail(email string) (model.User, error) {
	var user model.User
	row := u.db.QueryRow(`SELECT id, email, password FROM users`)
	err := row.Scan(&user.Id,
		&user.Email,
		&user.Password,
	)
	return user, err
}
