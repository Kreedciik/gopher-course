package service

import (
	"log/slog"
	"reservation/model"
	"reservation/pkg/repository"
)

type Reservation interface {
	CreateReservation(model.CreateReservationDTO) error
	UpdateReservation(model.UpdateReservationDTO) error
	RemoveReservation(string) error
	GetAllReservations() ([]model.Reservation, error)
	GetReservationById(string) (model.Reservation, error)
}

type ReservationService struct {
	repository *repository.ReservationRepository
}

func NewReservationService(repository *repository.ReservationRepository) *ReservationService {
	return &ReservationService{
		repository,
	}
}

func (r *ReservationService) CreateReservation(reservation model.CreateReservationDTO) error {
	if err := r.repository.InsertReservation(reservation); err != nil {
		slog.Error("create-reservation: ", err.Error())
		return err
	}
	return nil
}

func (r *ReservationService) UpdateReservation(reservation model.UpdateReservationDTO) error {
	if err := r.repository.UpdateReservation(reservation); err != nil {
		slog.Error("update-reservation: ", err.Error())
		return err
	}
	return nil
}

func (r *ReservationService) RemoveReservation(id string) error {
	if err := r.repository.RemoveReservation(id); err != nil {
		slog.Error("remove-reservation: ", err.Error())
		return err
	}
	return nil
}

func (r *ReservationService) GetAllReservations() ([]model.Reservation, error) {
	reservations, err := r.repository.FindAllReservations()
	if err != nil {
		slog.Error("get-all-reservations: ", err.Error())
		return reservations, err
	}
	return reservations, nil
}

func (r *ReservationService) GetReservationById(id string) (model.Reservation, error) {
	reservation, err := r.repository.FindReservationById(id)
	if err != nil {
		slog.Error("get-reservation-by-id: ", err.Error())
		return reservation, err
	}
	return reservation, nil
}
