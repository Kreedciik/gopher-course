package service

import (
	"farmish/model"
	"farmish/pkg/pubsub"
	"farmish/pkg/repository"
	"fmt"
	"log/slog"
	"time"
)

type Animal interface {
	CreateAnimal(model.CreateAnimalDTO) error
	UpdateAnimal(model.UpdateAnimalDTO) error
	DeleteAnimal(string) error
	GetAnimalById(string) (model.Animal, error)
	ToggleHunger(string) error
	ToggleHealth(string) error
	MonitorHealth(time.Duration)
	MonitorFeed(time.Duration)
}

type AnimalService struct {
	repository *repository.AnimalRepository
	pubSub     *pubsub.PubSub
}

func NewAnimalService(repository *repository.AnimalRepository, ps *pubsub.PubSub) *AnimalService {
	return &AnimalService{repository, ps}
}

func (a *AnimalService) CreateAnimal(newAnimal model.CreateAnimalDTO) error {
	if err := a.repository.InsertAnimal(newAnimal); err != nil {
		slog.Error(fmt.Sprintf("create-animal: %s", err.Error()))
		return fmt.Errorf("error while creating animal")
	}
	return nil
}

func (a *AnimalService) UpdateAnimal(animal model.UpdateAnimalDTO) error {
	if err := a.repository.UpdateAnimal(animal); err != nil {
		slog.Error(fmt.Sprintf("update-animal: %s", err.Error()))
		return fmt.Errorf("error while creating animal")
	}
	return nil
}

func (a *AnimalService) DeleteAnimal(id string) error {
	if err := a.repository.DeleteAnimal(id); err != nil {
		slog.Error(fmt.Sprintf("delete-animal: %s", err.Error()))
		return fmt.Errorf("error while deleting animal")
	}
	return nil
}

func (a *AnimalService) GetAnimalById(id string) (model.Animal, error) {
	animal, err := a.repository.FindAnimalById(id)
	if err != nil {
		slog.Error(fmt.Sprintf("update-animal: %s", err.Error()))
		return animal, fmt.Errorf("error while creating animal")
	}
	return animal, nil
}

func (a *AnimalService) ToggleHunger(id string) error {
	return a.repository.ToggleHungryAnimal(id)
}
func (a *AnimalService) ToggleHealth(id string) error {
	return a.repository.ToggleSickAnimal(id)
}

func (a *AnimalService) MonitorHealth(timeout time.Duration) {
	ticker := time.NewTicker(1 * time.Minute)

	for range ticker.C {
		animals, _ := a.repository.FindSickAnimals()
		now := time.Now()

		for _, animal := range animals {
			if now.Sub(animal.UpdatedAt) > timeout {
				message := fmt.Sprintf("ALERT: Animal %s is unhealthy for over %d minutes!", animal.Name, int(timeout.Minutes()))
				a.pubSub.Publish("notification", message)
			}
		}
	}
}

func (a *AnimalService) MonitorFeed(timeout time.Duration) {
	ticker := time.NewTicker(1 * time.Minute)

	for range ticker.C {
		animals, _ := a.repository.FindHungryAnimals()
		now := time.Now()

		for _, animal := range animals {
			if now.Sub(animal.UpdatedAt) > timeout {
				message := fmt.Sprintf("ALERT: Animal %s is hungry for over %d minutes!", animal.Name, int(timeout.Minutes()))
				a.pubSub.Publish("notification", message)
			}
		}
	}
}
