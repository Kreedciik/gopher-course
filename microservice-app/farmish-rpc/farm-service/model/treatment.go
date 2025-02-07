package model

import "time"

type CreateTreatmentDTO struct {
	Animal        string    `json:"animal" binding:"required"`
	Medicine      string    `json:"medicine" binding:"required"`
	Quantity      float64   `json:"quantity" binding:"required"`
	TreatmentDate time.Time `json:"treatmentDate" binding:"required"`
}
