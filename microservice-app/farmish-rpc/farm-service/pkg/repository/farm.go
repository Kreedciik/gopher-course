package repository

import (
	"database/sql"
	"farmish/model"
	"time"

	"github.com/google/uuid"
)

type Farm interface {
	InsertFarm(model.CreateFarmDTO) error
	UpdateFarm(model.UpdateFarmDTO) error
	FindAllFarm() ([]model.Farm, error)
	FindFarmById(string) (model.Farm, error)
	DeleteFarm(string) error
}

type FarmRepository struct {
	db *sql.DB
}

func NewFarmRepository(db *sql.DB) *FarmRepository {
	return &FarmRepository{db}
}

func (f *FarmRepository) InsertFarm(newFarm model.CreateFarmDTO) error {
	_, err := f.db.Exec(
		`INSERT INTO farms VALUES ($1, $2, $3, $4)`,
		uuid.NewString(),
		newFarm.Name,
		newFarm.Lat,
		newFarm.Lng,
	)
	return err
}
func (f *FarmRepository) UpdateFarm(farm model.UpdateFarmDTO) error {
	_, err := f.db.Exec(
		`UPDATE farms SET 
		name = $1, lat = $2,
		long = $3, updated_at = $4
		WHERE id = $5
		`,
		farm.Name,
		farm.Lat,
		farm.Lng,
		time.Now(),
		farm.Id,
	)
	return err
}
func (f *FarmRepository) FindAllFarm() ([]model.Farm, error) {
	farms := []model.Farm{}

	rows, err := f.db.Query(`SELECT (id, name, lat, long) FROM farms`)
	if err != nil {
		return farms, err
	}
	defer rows.Close()
	for rows.Next() {
		var farm model.Farm
		if err := rows.Scan(
			&farm.Id,
			&farm.Name,
			&farm.Lat,
			&farm.Lng,
		); err != nil {
			return farms, err
		}
		farms = append(farms, farm)
	}
	return farms, nil
}
func (f *FarmRepository) FindFarmById(id string) (model.Farm, error) {
	var farm model.Farm
	row := f.db.QueryRow(`
	SELECT id, name, lat, long FROM farms
	WHERE id = $1`, id)

	if err := row.Scan(
		&farm.Id,
		&farm.Name,
		&farm.Lat,
		&farm.Lng,
	); err != nil {
		return farm, err
	}
	return farm, nil
}
func (f *FarmRepository) DeleteFarm(id string) error {
	_, err := f.db.Exec(`DELETE FROM farms WHERE id = $1`, id)
	return err
}
