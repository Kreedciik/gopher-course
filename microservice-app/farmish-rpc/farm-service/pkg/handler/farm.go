package handler

import (
	"farmish/model"
)

func (h *Handler) CreateFarm(newFarm model.CreateFarmDTO, response *string) error {
	if err := h.services.Farm.CreateFarm(newFarm); err != nil {
		return err
	}
	*response = "farm created successfully"
	return nil
}

func (h *Handler) UpdateFarm(farm model.UpdateFarmDTO, response *string) error {
	if err := h.services.Farm.UpdateFarm(farm); err != nil {
		return err
	}
	*response = "farm updated successfully"
	return nil
}

func (h *Handler) GetFarm(id string, response *model.Farm) error {
	farm, err := h.services.Farm.GetFarmById(id)
	if err != nil {
		return err
	}
	*response = farm
	return nil
}

func (h *Handler) DeleteFarm(id string, response *string) error {
	if err := h.services.Farm.DeleteFarm(id); err != nil {
		return err
	}
	*response = "deleted successfully"
	return nil
}
