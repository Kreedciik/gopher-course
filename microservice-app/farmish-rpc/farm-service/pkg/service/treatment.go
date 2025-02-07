package service

import (
	"farmish/model"
	"farmish/pkg/repository"
	"fmt"
	"log/slog"
)

type Treatment interface {
	TreatAnimal(model.CreateTreatmentDTO) error
}

type TreatmentService struct {
	repository          *repository.TreatmentRepository
	warehouseRepository *repository.WarehouseRepository
}

func NewTreatmentService(repository *repository.TreatmentRepository, warehouseRepository *repository.WarehouseRepository) *TreatmentService {
	return &TreatmentService{
		repository,
		warehouseRepository,
	}
}

func (f *TreatmentService) TreatAnimal(newTreatment model.CreateTreatmentDTO) error {
	stock, err := f.warehouseRepository.FindStockByName(newTreatment.Medicine)
	if err != nil {
		slog.Error(fmt.Sprintf("treat-animal: %s", err.Error()))
		return fmt.Errorf("the medicine %s is not registered", newTreatment.Medicine)
	}

	if stock.Quantity < newTreatment.Quantity {
		slog.Error(fmt.Sprintf("feed-animal: %s", "not enough quantity"))
		return fmt.Errorf("no enough quantity of feed %s", newTreatment.Medicine)
	}

	return nil
}
