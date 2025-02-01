package service

import (
	"auth/pkg/repository"
)

type Service struct {
	User
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repository.User),
	}
}
