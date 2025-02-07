package model

import "time"

type CreateStockDTO struct {
	Name      string  `json:"name" binding:"required"`
	Quantity  float64 `json:"quantity"`
	Cost      float64 `json:"cost"`
	StockType string  `json:"stockType" binding:"required"`
}

type Stock struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	StockType string  `json:"stockType"`
	Quantity  float64 `json:"quantity"`
	Cost      float64 `json:"cost"`
}

type SupplyStockDTO struct {
	Type        string    `json:"type" binding:"required"`
	Quantity    float64   `json:"quantity" binding:"required"`
	ArrivalDate time.Time `json:"arrivalDate" binding:"required"`
	Notes       string    `json:"notes"`
}
