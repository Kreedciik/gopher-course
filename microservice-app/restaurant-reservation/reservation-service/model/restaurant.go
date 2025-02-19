package model

type Restaurant struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	Description string `json:"description"`
}
type CreateRestaurantDTO struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	Description string `json:"description"`
}
type UpdateRestaurantDTO = Restaurant
