package service

import (
	"farmish/model"
	"farmish/pkg/repository"
	"fmt"
	"log/slog"
)

type Warehouse interface {
	CreateStock(model.CreateStockDTO) error
	SupplyFeedToWarehouse(model.SupplyStockDTO) error
	SupplyMedicineToWarehouse(model.SupplyStockDTO) error
}

type WarehouseService struct {
	repository *repository.WarehouseRepository
}

func NewWarehouseService(repository *repository.WarehouseRepository) *WarehouseService {
	return &WarehouseService{
		repository,
	}
}

func (w *WarehouseService) CreateStock(newStock model.CreateStockDTO) error {
	stock, _ := w.repository.FindStockByName(newStock.Name)
	if stock.Id != "" {
		slog.Error(fmt.Sprintf("create-stock: %s", "duplicated type value"))
		return fmt.Errorf("stock type should be unique")
	}
	return w.repository.InsertStock(newStock)
}

func (w *WarehouseService) SupplyFeedToWarehouse(stock model.SupplyStockDTO) error {
	err := w.repository.UpdateFeedStock(stock)
	if err != nil {
		slog.Error(fmt.Sprintf("supply-feed: %s", err.Error()))
		return fmt.Errorf("failed to supply feed")
	}
	return nil
}
func (w *WarehouseService) SupplyMedicineToWarehouse(stock model.SupplyStockDTO) error {
	err := w.repository.UpdateMedicineStock(stock)
	if err != nil {
		slog.Error(fmt.Sprintf("supply-medicine: %s", err.Error()))
		return fmt.Errorf("failed to supply medicine")
	}
	return nil
}
