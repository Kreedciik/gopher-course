package service

import (
	"farmish/pkg/pubsub"
	"farmish/pkg/repository"
)

type Service struct {
	User
	Farm
	Animal
	Warehouse
	Feeding
	Treatment
	Dashboard
}

func NewService(repository *repository.Repository, ps *pubsub.PubSub) *Service {
	return &Service{
		User:      NewUserService(repository.User),
		Farm:      NewFarmService(repository.Farm),
		Animal:    NewAnimalService(repository.Animal, ps),
		Warehouse: NewWarehouseService(repository.Warehouse),
		Feeding:   NewFeedingService(repository.Feeding, repository.Warehouse),
		Treatment: NewTreatmentService(repository.Treatment, repository.Warehouse),
		Dashboard: NewDashboardService(repository.Animal, repository.Warehouse),
	}
}
