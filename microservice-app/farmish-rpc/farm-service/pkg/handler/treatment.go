package handler

import (
	"farmish/model"
)

func (h *Handler) Treatment(treatment model.CreateTreatmentDTO, response *string) error {
	if err := h.services.Treatment.TreatAnimal(treatment); err != nil {
		return err
	}
	*response = "success"
	return nil
}
