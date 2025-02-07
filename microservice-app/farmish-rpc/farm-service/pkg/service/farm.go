package service

import (
	"farmish/model"
	"farmish/pkg/repository"
	"fmt"
	"log/slog"
)

type Farm interface {
	CreateFarm(model.CreateFarmDTO) error
	UpdateFarm(model.UpdateFarmDTO) error
	GetAllFarm() ([]model.Farm, error)
	GetFarmById(string) (model.Farm, error)
	DeleteFarm(string) error
}

type FarmService struct {
	repository *repository.FarmRepository
}

func NewFarmService(repository *repository.FarmRepository) *FarmService {
	return &FarmService{repository}
}

func (f *FarmService) CreateFarm(newFarm model.CreateFarmDTO) error {
	if err := f.repository.InsertFarm(newFarm); err != nil {
		slog.Error(fmt.Sprintf("create-farm: %s", err.Error()))
		return fmt.Errorf("error while creating farm")
	}
	return nil
}
func (f *FarmService) UpdateFarm(farm model.UpdateFarmDTO) error {
	if err := f.repository.UpdateFarm(farm); err != nil {
		slog.Error(fmt.Sprintf("update-farm: %s", err.Error()))
		return fmt.Errorf("error while updating farm")
	}
	return nil
}
func (f *FarmService) GetAllFarm() ([]model.Farm, error) {
	farms, err := f.repository.FindAllFarm()
	if err != nil {
		slog.Error(fmt.Sprintf("get-farms: %s", err.Error()))
		return farms, err
	}
	return farms, nil
}
func (f *FarmService) GetFarmById(id string) (model.Farm, error) {
	farm, err := f.repository.FindFarmById(id)
	if err != nil {
		slog.Error(fmt.Sprintf("get-farm-by-id: %s", err.Error()))
		return farm, err
	}
	if farm.Id == "" {
		return farm, fmt.Errorf("not found")
	}
	return farm, nil
}
func (f *FarmService) DeleteFarm(id string) error {
	if err := f.repository.DeleteFarm(id); err != nil {
		slog.Error(fmt.Sprintf("delete-farm: %s", err.Error()))
		return err
	}
	return nil
}
