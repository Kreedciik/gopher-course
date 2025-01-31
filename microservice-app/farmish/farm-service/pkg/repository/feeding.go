package repository

import (
	"database/sql"
	"farmish/model"

	"github.com/google/uuid"
)

type Feeding interface {
	InsertFeeding(model.CreateFeedingDTO) error
}

type FeedingRepository struct {
	db *sql.DB
}

func NewFeedingRepository(db *sql.DB) *FeedingRepository {
	return &FeedingRepository{db}
}

func (f *FeedingRepository) InsertFeeding(newFeeding model.CreateFeedingDTO) error {
	tx, err := f.db.Begin()
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
		`INSERT INTO feeding_record VALUES ($1, $2, $3, $4, $5, $6)`,
		uuid.NewString(),
		newFeeding.Type,
		newFeeding.Quantity,
		newFeeding.FeedingTime,
		newFeeding.Notes,
		newFeeding.Animal,
	)

	if err != nil {
		return err
	}

	_, err = tx.Exec(
		`UPDATE warehouse SET quantity = quantity - $1 WHERE type = $2`,
		newFeeding.Quantity,
		newFeeding.Type,
	)

	return err
}
