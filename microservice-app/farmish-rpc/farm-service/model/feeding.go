package model

import "time"

type CreateFeedingDTO struct {
	Type        string    `json:"type" binding:"required"`
	Quantity    float64   `json:"quantity" binding:"required"`
	FeedingTime time.Time `json:"feedingTime" binding:"required"`
	Notes       string    `json:"notes"`
	Animal      string    `json:"animal" binding:"required"`
}
