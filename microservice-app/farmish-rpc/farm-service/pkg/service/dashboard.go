package service

import (
	"farmish/model"
	"farmish/pkg/repository"
	"fmt"
	"log/slog"
)

type Dashboard interface {
	GetDashboardData() (model.Dashboard, error)
}

type DashboardService struct {
	animalRepository    *repository.AnimalRepository
	warehouseRepository *repository.WarehouseRepository
}

func NewDashboardService(
	animalRepository *repository.AnimalRepository,
	warehouseRepository *repository.WarehouseRepository) *DashboardService {
	return &DashboardService{
		animalRepository,
		warehouseRepository,
	}
}

func (d *DashboardService) GetDashboardData() (model.Dashboard, error) {
	var dashboard model.Dashboard
	stat, err := d.animalRepository.FindAnimalsQuantity()
	if err != nil {
		slog.Error(fmt.Sprintf("dashboard: %s", err.Error()))
		return dashboard, fmt.Errorf("could not get quantity and average weight")
	}
	dashboard.Quantity = stat.Quantity
	dashboard.AverageWeight = stat.AverageWeight

	sickAnimals, err := d.animalRepository.FindSickAnimals()
	if err != nil {
		slog.Error(fmt.Sprintf("dashboard: %s", err.Error()))
		return dashboard, fmt.Errorf("could not get sick animals")
	}
	dashboard.SickAnimals = sickAnimals

	hungryAnimals, err := d.animalRepository.FindHungryAnimals()
	if err != nil {
		slog.Error(fmt.Sprintf("dashboard: %s", err.Error()))
		return dashboard, fmt.Errorf("could not get hungry animals")
	}
	dashboard.HungryAnimals = hungryAnimals

	allFeed, err := d.warehouseRepository.FindFeed()
	if err != nil {
		slog.Error(fmt.Sprintf("dashboard: %s", err.Error()))
		return dashboard, fmt.Errorf("could not get all feed")
	}
	dashboard.Food = allFeed

	allMedicine, err := d.warehouseRepository.FindMedicine()
	if err != nil {
		slog.Error(fmt.Sprintf("dashboard: %s", err.Error()))
		return dashboard, fmt.Errorf("could not get all medicine")
	}
	dashboard.Medicine = allMedicine

	return dashboard, nil
}
