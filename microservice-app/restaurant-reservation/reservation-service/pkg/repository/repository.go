package repository

import "database/sql"

type Repository struct {
	Restaurant  *RestaurantRepository
	Reservation *ReservationRepository
}

func NewRepository(postgresDB *sql.DB) *Repository {
	return &Repository{
		Restaurant:  NewRestaurantRepository(postgresDB),
		Reservation: NewReservationRepository(postgresDB),
	}
}
