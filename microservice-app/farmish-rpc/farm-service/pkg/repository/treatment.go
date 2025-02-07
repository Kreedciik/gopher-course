package repository

import (
	"database/sql"
	"farmish/model"

	"github.com/google/uuid"
)

type Treatment interface {
	InsertTreatment(model.CreateTreatmentDTO) error
}

type TreatmentRepository struct {
	db *sql.DB
}

func NewTreatmentRepository(db *sql.DB) *TreatmentRepository {
	return &TreatmentRepository{db}
}

func (t *TreatmentRepository) InsertTreatment(newTreatment model.CreateTreatmentDTO) error {
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	_, err = tx.Exec(
		`INSERT INTO treatments VALUES ($1, $2, $3, $4, $5)`,
		uuid.NewString(),
		newTreatment.Animal,
		newTreatment.Medicine,
		newTreatment.Quantity,
		newTreatment.TreatmentDate,
	)

	if err != nil {
		return err
	}

	_, err = tx.Exec(
		`UPDATE warehouse SET quantity = quantity - $1 WHERE type = $2`,
		newTreatment.Quantity,
		newTreatment.Medicine,
	)

	return err
}
