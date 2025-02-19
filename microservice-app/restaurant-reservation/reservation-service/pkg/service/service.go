package service

import (
	"reservation/pkg/repository"
)

type Service struct {
	Restaurant
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Restaurant: NewRestaurantService(repository.Restaurant),
	}
}
