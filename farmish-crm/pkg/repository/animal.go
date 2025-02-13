package repository

import (
	"database/sql"
	"farmish/model"

	"github.com/google/uuid"
)

type Animal interface {
	InsertAnimal(model.CreateAnimalDTO) error
	ToggleSickAnimal(string) error
	ToggleHungryAnimal(string) error
	UpdateAnimal(model.UpdateAnimalDTO) error
	DeleteAnimal(string) error
	FindAnimalById(string) (model.Animal, error)
	FindAnimalsQuantity() (model.AnimalQuantityStatistics, error)
	FindSickAnimals() ([]model.Animal, error)
	FindHungryAnimals() ([]model.Animal, error)
}

type AnimalRepository struct {
	db *sql.DB
}

func NewAnimalRepository(db *sql.DB) *AnimalRepository {
	return &AnimalRepository{db}
}

func (a *AnimalRepository) ToggleSickAnimal(id string) error {
	_, err := a.db.Exec(
		`UPDATE animals SET is_healthy = NOT is_healthy WHERE id = $1`,
		id,
	)
	return err
}
func (a *AnimalRepository) ToggleHungryAnimal(id string) error {
	_, err := a.db.Exec(
		`UPDATE animals SET is_hungry = NOT is_hungry WHERE id = $1`,
		id,
	)
	return err
}

func (a *AnimalRepository) InsertAnimal(newAnimal model.CreateAnimalDTO) error {
	_, err := a.db.Exec(
		`INSERT INTO animals
		(id, name, species, breed, birthdate, gender, weight, farm)
		VALUES
		($1, $2, $3, $4, $5, $6, $7, $8)`,
		uuid.NewString(),
		newAnimal.Name,
		newAnimal.Species,
		newAnimal.Breed,
		newAnimal.Birthdate,
		newAnimal.Gender,
		newAnimal.Weight,
		newAnimal.Farm,
	)
	return err
}

func (a *AnimalRepository) UpdateAnimal(animal model.UpdateAnimalDTO) error {
	_, err := a.db.Exec(
		`UPDATE animals SET
		name = $1, species = $2, breed = $3, birthdate = $4, 
		gender = $5, weight = $6, farm = $7
		WHERE id = $8
		`,
		animal.Name,
		animal.Species,
		animal.Breed,
		animal.Birthdate,
		animal.Gender,
		animal.Weight,
		animal.Farm,
		animal.Id,
	)
	return err
}

func (a *AnimalRepository) DeleteAnimal(id string) error {
	_, err := a.db.Exec(`DELETE FROM animals WHERE id = $1`, id)
	return err
}

func (a *AnimalRepository) FindSickAnimals() ([]model.Animal, error) {
	var sickAnimals []model.Animal
	rows, err := a.db.Query(
		`SELECT id, name, species, 
		breed, birthdate, gender, 
		weight, is_healthy, is_hungry,
		updated_at
		farm FROM animals WHERE is_healthy = false`,
	)
	if err != nil {
		return sickAnimals, err
	}
	defer rows.Close()

	for rows.Next() {
		var animal model.Animal
		if err := rows.Scan(
			&animal.Id,
			&animal.Name,
			&animal.Species,
			&animal.Breed,
			&animal.Birthdate,
			&animal.Gender,
			&animal.Weight,
			&animal.IsHealthy,
			&animal.IsHungry,
			&animal.UpdatedAt,
		); err != nil {
			return sickAnimals, err
		}
		sickAnimals = append(sickAnimals, animal)
	}

	return sickAnimals, nil
}

func (a *AnimalRepository) FindHungryAnimals() ([]model.Animal, error) {
	var hungryAnimals []model.Animal
	rows, err := a.db.Query(
		`SELECT id, name, species, 
		breed, birthdate, gender, 
		weight, is_healthy, is_hungry, updated_at FROM animals WHERE is_hungry = true`,
	)
	if err != nil {
		return hungryAnimals, err
	}
	defer rows.Close()

	for rows.Next() {
		var animal model.Animal
		if err := rows.Scan(
			&animal.Id,
			&animal.Name,
			&animal.Species,
			&animal.Breed,
			&animal.Birthdate,
			&animal.Gender,
			&animal.Weight,
			&animal.IsHealthy,
			&animal.IsHungry,
			&animal.UpdatedAt,
		); err != nil {
			return hungryAnimals, err
		}
		hungryAnimals = append(hungryAnimals, animal)
	}

	return hungryAnimals, nil
}

func (a *AnimalRepository) FindAnimalsQuantity() (model.AnimalQuantityStatistics, error) {
	var statistics model.AnimalQuantityStatistics
	row := a.db.QueryRow(
		`SELECT AVG(weight), COUNT(*) FROM animals`,
	)

	err := row.Scan(
		&statistics.AverageWeight,
		&statistics.Quantity,
	)

	return statistics, err
}

func (a *AnimalRepository) FindAnimalById(id string) (model.Animal, error) {
	var animal model.Animal

	row := a.db.QueryRow(
		`SELECT id, name, species, breed, birthdate,
		gender, weight, isHealthy, farm FROM animals
		WHERE id = $1`,
		id,
	)

	if err := row.Scan(
		&animal.Id,
		&animal.Name,
		&animal.Species,
		&animal.Breed,
		&animal.Birthdate,
		&animal.Gender,
		&animal.Weight,
		&animal.IsHealthy,
		&animal.Farm,
	); err != nil {
		return animal, err
	}

	return animal, nil
}
