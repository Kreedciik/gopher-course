package model

import "time"

type Reservation struct {
	Id              string    `json:"id"`
	UserId          string    `json:"userId"`
	RestaurantId    string    `json:"restaurantId"`
	ReservationTime time.Time `json:"reservationTime"`
	Status          string    `json:"status"`
}
type CreateReservationDTO struct {
	UserId          string    `json:"userId"`
	RestaurantId    string    `json:"restaurantId"`
	ReservationTime time.Time `json:"reservationTime"`
	Status          string    `json:"status"`
}
type UpdateReservationDTO = Reservation
