package service

import (
	"log/slog"
	"reservation/model"
	"reservation/pkg/repository"
)

type Restaurant interface {
	CreateRestaurant(model.CreateRestaurantDTO) error
	UpdateRestaurant(model.UpdateRestaurantDTO) error
	RemoveRestaurant(string) error
	GetAllRestaurants() ([]model.Restaurant, error)
	GetRestaurantById(string) (model.Restaurant, error)
}

type RestaurantService struct {
	repository *repository.RestaurantRepository
}

func NewRestaurantService(repository *repository.RestaurantRepository) *RestaurantService {
	return &RestaurantService{
		repository,
	}
}

func (r *RestaurantService) CreateRestaurant(restaurant model.CreateRestaurantDTO) error {
	if err := r.repository.InsertRestaurant(restaurant); err != nil {
		slog.Error("create-restaurant: ", err.Error())
		return err
	}
	return nil
}

func (r *RestaurantService) UpdateRestaurant(restaurant model.UpdateRestaurantDTO) error {
	if err := r.repository.UpdateRestaurant(restaurant); err != nil {
		slog.Error("update-restaurant: ", err.Error())
		return err
	}
	return nil
}

func (r *RestaurantService) RemoveRestaurant(id string) error {
	if err := r.repository.RemoveRestaurant(id); err != nil {
		slog.Error("remove-restaurant: ", err.Error())
		return err
	}
	return nil
}

func (r *RestaurantService) GetAllRestaurants() ([]model.Restaurant, error) {
	restaurants, err := r.repository.FindAllRestaurants()
	if err != nil {
		slog.Error("get-all-restaurants: ", err.Error())
		return restaurants, err
	}
	return restaurants, nil
}

func (r *RestaurantService) GetRestaurantById(id string) (model.Restaurant, error) {
	restaurant, err := r.repository.FindRestaurantById(id)
	if err != nil {
		slog.Error("get-all-restaurants: ", err.Error())
		return restaurant, err
	}
	return restaurant, nil
}
