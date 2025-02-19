package repository

import (
	"database/sql"
	"reservation/model"

	"github.com/google/uuid"
)

type Reservation interface {
	InsertReservation(model.CreateReservationDTO) error
	UpdateReservation(model.UpdateReservationDTO) error
	FindAllReservations() ([]model.Reservation, error)
	FindReservationById(string) (model.Reservation, error)
	RemoveReservation(string) error
}
type ReservationRepository struct {
	db *sql.DB
}

func NewReservationRepository(db *sql.DB) *ReservationRepository {
	return &ReservationRepository{db}
}

func (r *ReservationRepository) InsertReservation(reservation model.CreateReservationDTO) error {
	_, err := r.db.Exec(`
	INSERT INTO reservations VALUES ($1, $2, $3, $4, $5)
	`,
		uuid.NewString(),
		reservation.UserId,
		reservation.RestaurantId,
		reservation.ReservationTime,
		reservation.Status,
	)
	return err
}

func (r *ReservationRepository) UpdateReservation(reservation model.UpdateReservationDTO) error {
	_, err := r.db.Exec(`
	UPDATE reservations SET 
	user_id = $1,
	restaurant_id = $2,
	reservation_time = $3,
	status = $4
	WHERE id = %5
	`,
		reservation.UserId,
		reservation.RestaurantId,
		reservation.ReservationTime,
		reservation.Status,
		reservation.Id,
	)
	return err
}

func (r *ReservationRepository) RemoveReservation(id string) error {
	_, err := r.db.Exec(`
	DELETE FROM reservations WHERE id = $1
	`,
		id,
	)
	return err
}

func (r *ReservationRepository) FindReservationById(id string) (model.Reservation, error) {
	var reservation model.Reservation
	row := r.db.QueryRow(`SELECT id, user_id, restaurant_id, reservation_time, status FROM reservations WHERE id = $1`, id)
	err := row.Scan(
		&reservation.Id,
		&reservation.UserId,
		&reservation.RestaurantId,
		&reservation.ReservationTime,
		&reservation.Status,
	)
	return reservation, err
}

func (r *ReservationRepository) FindAllReservations() ([]model.Reservation, error) {
	reservations := []model.Reservation{}

	rows, err := r.db.Query(`SELECT id, user_id, restaurant_id, reservation_time, status FROM reservations`)
	if err != nil {
		return reservations, err
	}

	for rows.Next() {
		var reservation model.Reservation
		err = rows.Scan(
			&reservation.Id,
			&reservation.UserId,
			&reservation.RestaurantId,
			&reservation.ReservationTime,
			&reservation.Status,
		)
		if err != nil {
			return reservations, err
		}
		reservations = append(reservations, reservation)
	}

	return reservations, nil
}
