package repository

import (
	"database/sql"
	"reservation/model"

	"github.com/google/uuid"
)

type Restaurant interface {
	InsertRestaurant(model.CreateRestaurantDTO) error
	UpdateRestaurant(model.UpdateRestaurantDTO) error
	RemoveRestaurant(string) error
	FindAllRestaurants() ([]model.Restaurant, error)
	FindRestaurantById(string) (model.Restaurant, error)
}
type RestaurantRepository struct {
	db *sql.DB
}

func NewRestaurantRepository(db *sql.DB) *RestaurantRepository {
	return &RestaurantRepository{db}
}

func (r *RestaurantRepository) InsertRestaurant(newRestaurant model.CreateRestaurantDTO) error {
	_, err := r.db.Exec(`
	INSERT INTO restaurants VALUES ($1, $2, $3, $4, $5)
	`,
		uuid.NewString(),
		newRestaurant.Name,
		newRestaurant.Address,
		newRestaurant.PhoneNumber,
		newRestaurant.Description,
	)
	return err
}

func (r *RestaurantRepository) UpdateRestaurant(restaurant model.UpdateRestaurantDTO) error {
	_, err := r.db.Exec(`
	UPDATE restaurants SET 
	name = $1,
	address = $2,
	phone_number = $3,
	description = $4
	WHERE id = %5
	`,
		restaurant.Name,
		restaurant.Address,
		restaurant.PhoneNumber,
		restaurant.Description,
		restaurant.Id,
	)
	return err
}

func (r *RestaurantRepository) RemoveRestaurant(id string) error {
	_, err := r.db.Exec(`
	DELETE FROM restaurants WHERE id = $1
	`,
		id,
	)
	return err
}

func (r *RestaurantRepository) FindRestaurantById(id string) (model.Restaurant, error) {
	var restaurant model.Restaurant
	row := r.db.QueryRow(`SELECT id, name, address, phone_number, description FROM restaurants`)
	err := row.Scan(
		&restaurant.Id,
		&restaurant.Name,
		&restaurant.Address,
		&restaurant.PhoneNumber,
		&restaurant.Description,
	)
	return restaurant, err
}

func (r *RestaurantRepository) FindAllRestaurants() ([]model.Restaurant, error) {
	restaurants := []model.Restaurant{}

	rows, err := r.db.Query(`SELECT id, name, address, phone_number, description FROM restaurants`)
	if err != nil {
		return restaurants, err
	}

	for rows.Next() {
		var restaurant model.Restaurant
		err = rows.Scan(
			&restaurant.Id,
			&restaurant.Name,
			&restaurant.Address,
			&restaurant.PhoneNumber,
			&restaurant.Description,
		)
		if err != nil {
			return restaurants, err
		}
		restaurants = append(restaurants, restaurant)
	}

	return restaurants, nil
}
