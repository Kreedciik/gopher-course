package handler

import (
	"farmish/model"
)

func (h *Handler) Feeding(feeding model.CreateFeedingDTO, response *string) error {

	if err := h.services.Feeding.FeedAnimal(feeding); err != nil {
		return err
	}
	*response = "success"
	return nil
}
