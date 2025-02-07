package handler

import (
	"errors"
	"farmish/model"
)

func (h *Handler) CreateStock(newStock model.CreateStockDTO, response *string) error {
	if newStock.Cost < 0 || newStock.Quantity < 0 {
		return errors.New("value should not be negative")
	}

	if err := h.services.Warehouse.CreateStock(newStock); err != nil {
		return err
	}
	*response = "stock created successfully"
	return nil
}

func (h *Handler) SupplyFeedStock(stock model.SupplyStockDTO, response *string) error {
	if stock.Quantity < 0 {
		return errors.New("value should not be negative")
	}

	if err := h.services.Warehouse.SupplyFeedToWarehouse(stock); err != nil {
		return err
	}
	*response = "stock supplied successfully"
	return nil
}

func (h *Handler) SupplyMedicineStock(stock model.SupplyStockDTO, response *string) error {
	if stock.Quantity < 0 {
		return errors.New("value should not be negative")
	}

	if err := h.services.Warehouse.SupplyMedicineToWarehouse(stock); err != nil {
		return err
	}
	*response = "stock supplied successfully"
	return nil
}
