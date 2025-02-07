package model

import "time"

type CreateAnimalDTO struct {
	Name      string  `json:"name" binding:"required"`
	Species   string  `json:"species" binding:"required"`
	Breed     string  `json:"breed" binding:"required"`
	Birthdate string  `json:"birthdate" binding:"required"`
	Gender    string  `json:"gender" binding:"required"`
	Weight    float64 `json:"weight" binding:"required"`
	Farm      string  `json:"farm" binding:"required"`
}

type Animal struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Species   string    `json:"species"`
	Breed     string    `json:"breed"`
	Birthdate string    `json:"birthdate"`
	Gender    string    `json:"gender"`
	Weight    float64   `json:"weight"`
	IsHealthy bool      `json:"isHealthy"`
	IsHungry  bool      `json:"isHungry"`
	Farm      string    `json:"farm"`
	UpdatedAt time.Time `json:"-"`
}

type UpdateAnimalDTO = Animal

type AnimalQuantityStatistics struct {
	Quantity      int
	AverageWeight float64
}
