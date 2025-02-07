package handler

import (
	"farmish/model"
	"fmt"
)

func (h *Handler) CreateAnimal(newAnimal model.CreateAnimalDTO, response *string) error {
	if err := h.services.Animal.CreateAnimal(newAnimal); err != nil {
		return err
	}
	*response = "animal created successfully"
	return nil
}

func (h *Handler) UpdateAnimal(animal model.UpdateAnimalDTO, response *string) error {
	if err := h.services.Animal.UpdateAnimal(animal); err != nil {
		return err
	}
	*response = "animal updated successfully"
	return nil
}

func (h *Handler) DeleteAnimal(id string, response *string) error {
	if err := h.services.Animal.DeleteAnimal(id); err != nil {
		return err
	}
	*response = "deleted successfully"
	return nil
}

func (h *Handler) GetAnimalById(id string, response *model.Animal) error {
	animal, err := h.services.Animal.GetAnimalById(id)
	if err != nil {
		return err
	}

	*response = animal
	return nil
}

func (h *Handler) ToggleHealthyAnimal(id string, response *string) error {
	if err := h.services.Animal.ToggleHealth(id); err != nil {
		return err
	}
	*response = "success"
	return nil
}

func (h *Handler) ToggleHungryAnimal(id string, response *string) error {
	if err := h.services.Animal.ToggleHunger(id); err != nil {
		return err
	}
	*response = "success"
	return nil
}

func (h *Handler) ListenNotifications() {
	subscriber := h.ps.Subscribe("notifications")

	go func() {
		for message := range subscriber {
			// Websocket might be implemented
			fmt.Println("Notification Received:", message)
		}
	}()
}
