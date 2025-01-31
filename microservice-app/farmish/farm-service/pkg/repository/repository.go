package repository

import "database/sql"

type Repository struct {
	Farm      *FarmRepository
	Animal    *AnimalRepository
	Warehouse *WarehouseRepository
	Feeding   *FeedingRepository
	Treatment *TreatmentRepository
}

func NewRepository(postgresDB *sql.DB) *Repository {
	return &Repository{
		Farm:      NewFarmRepository(postgresDB),
		Animal:    NewAnimalRepository(postgresDB),
		Warehouse: NewWarehouseRepository(postgresDB),
		Feeding:   NewFeedingRepository(postgresDB),
		Treatment: NewTreatmentRepository(postgresDB),
	}
}
