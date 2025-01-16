package repository

import (
	"blogpost/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type User interface {
	InsertUser(models.UserCreateDTO) error
	InsertOneUser(models.UserCreateDTO) error
	InsertManyUser([]models.UserCreateDTO) error
	FindUserByEmail(string) (models.User, error)
	FindUserById(string) (models.User, error)
	UpdateUser(models.UpdateUserDTO) error
	FollowUser(userId string) error
}

type UserRepository struct {
	db             *sql.DB
	userCollection *mongo.Collection
}

func NewUserRepository(db *sql.DB, userCollection *mongo.Collection) *UserRepository {
	return &UserRepository{
		db,
		userCollection,
	}
}

var (
	usersTable     = "users"
	followersTable = "followers"
)

func (u *UserRepository) InsertOneUser(newUser models.UserCreateDTO) error {
	_, err := u.userCollection.InsertOne(context.TODO(), newUser)
	return err
}

func (u *UserRepository) InsertManyUser(users []models.UserCreateDTO) error {
	_, err := u.userCollection.InsertMany(context.TODO(), users)
	return err
}

func (u *UserRepository) InsertUser(newUser models.UserCreateDTO) error {

	query := fmt.Sprintf(
		`INSERT INTO %s 
	(id, name, email, password)
	VALUES ($1, $2, $3, $4)
	`, usersTable)

	_, err := u.db.Exec(query,
		uuid.NewString(),
		newUser.Name,
		newUser.Email,
		newUser.Password,
	)
	return err
}

func (u *UserRepository) FindUserByEmail(email string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf(`SELECT id, name, email, password FROM %s 
							WHERE email = $1`, usersTable)
	row := u.db.QueryRow(query, email)
	err := row.Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
	)
	return user, err
}

func (u *UserRepository) FindUserById(id string) (models.User, error) {
	var (
		user      models.User
		followers []models.Follower = []models.Follower{}
	)

	query := fmt.Sprintf(
		`SELECT u.id, u.name, u.email, u.password FROM %s u 
		WHERE id = $1`, usersTable)
	row := u.db.QueryRow(query, id)
	if err := row.Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
	); err != nil {
		return user, err
	}

	followersQuery := fmt.Sprintf(
		`SELECT f.user_id, u.name, u.email FROM %s f
		INNER JOIN %s u ON f.user_id = u.id
		`,
		followersTable,
		usersTable,
	)

	rows, err := u.db.Query(followersQuery)
	if err != nil {
		return user, err
	}
	defer rows.Close()
	for rows.Next() {
		var follower models.Follower
		if err := rows.Scan(
			&follower.Id,
			&follower.Name,
			&follower.Email,
		); err != nil {
			return user, err
		}

		followers = append(followers, follower)
	}

	user.Followers = followers
	return user, err
}

func (u *UserRepository) UpdateUser(user models.UpdateUserDTO) error {
	query := fmt.Sprintf(`UPDATE %s SET name = $1 WHERE id = $2`, usersTable)
	r, err := u.db.Exec(query, user.Name, user.Id)
	if err != nil {
		return err
	}
	if n, err := r.RowsAffected(); n == 0 || err != nil {
		return fmt.Errorf("user is not found")
	}

	return nil
}

func (u *UserRepository) FollowUser(userId, followerId string) error {
	var isSubscribed bool

	query := fmt.Sprintf(`SELECT EXISTS(
		SELECT 1 FROM %s WHERE follower_id = $1 AND user_id = $2
		)`, followersTable)

	row := u.db.QueryRow(query, followerId, userId)
	if err := row.Scan(&isSubscribed); err != nil {
		return err
	}

	if isSubscribed {
		deleteQuery := fmt.Sprintf(`DELETE FROM %s WHERE user_id = $1 AND follower_id = $2`, followersTable)
		_, err := u.db.Exec(deleteQuery, userId, followerId)
		if err != nil {
			return err
		}

		return nil
	}

	insertQuery := fmt.Sprintf(
		`INSERT INTO %s (id, user_id, follower_id) VALUES ($1, $2, $3)`,
		followersTable)

	_, err := u.db.Exec(insertQuery, uuid.NewString(), userId, followerId)
	return err
}
