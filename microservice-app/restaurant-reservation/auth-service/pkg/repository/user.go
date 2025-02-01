package repository

import (
	"auth/model"
	"database/sql"

	"github.com/google/uuid"
)

type User interface {
	InsertUser(model.CreateUserDTO) error
	FindUserByEmail(string) (model.User, error)
	FindUserById(string) (model.User, error)
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
		newUser.UserName,
		newUser.Password,
	)
	return err
}

func (u *UserRepository) FindUserByEmail(email string) (model.User, error) {
	var user model.User
	row := u.db.QueryRow(`SELECT id, email, password, user_name FROM users`)
	err := row.Scan(&user.Id,
		&user.Email,
		&user.Password,
		&user.Username,
	)
	return user, err
}

func (u *UserRepository) FindUserById(id string) (model.User, error) {
	var user model.User
	row := u.db.QueryRow(`SELECT id, user_name, email FROM users`)
	err := row.Scan(
		&user.Id,
		&user.Username,
		&user.Email,
	)
	return user, err
}
