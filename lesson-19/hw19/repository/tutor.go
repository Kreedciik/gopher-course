package repository

import (
	"database/sql"
	"hw19/model"
)

type tutorRepository struct {
	Db *sql.DB
}

func CreateTutorRepository(db *sql.DB) tutorRepository {
	return tutorRepository{Db: db}
}

func (t *tutorRepository) CreateTutor(tutor model.Tutor) error {
	_, err := t.Db.Exec(`INSERT INTO tutor (
		tutor_id, name, last_name, 
		email
		) VALUES 
		($1, $2, $3, $4)`,
		tutor.Id, tutor.Name,
		tutor.LastName, tutor.Email,
	)
	return err
}

func (t *tutorRepository) GetTutor(id string) (model.Tutor, error) {
	var tutor model.Tutor
	row := t.Db.QueryRow(`SELECT tutor_id, name, last_name, email
		FROM tutor
		WHERE tutor_id = $1
	`, id)

	err := row.Scan(
		&tutor.Id, &tutor.Name,
		&tutor.LastName, &tutor.Email,
	)

	return tutor, err
}

func (t *tutorRepository) GetAllTutors() ([]model.Tutor, error) {
	var tutors []model.Tutor

	rows, err := t.Db.Query(
		`SELECT tutor_id, name, last_name, email FROM tutor`)

	for rows.Next() {
		var tutor model.Tutor
		err = rows.Scan(
			&tutor.Id, &tutor.Name,
			&tutor.LastName, &tutor.Email,
		)
		tutors = append(tutors, tutor)
	}
	return tutors, err
}

func (t *tutorRepository) UpdateTutor(tutor model.Tutor) error {
	_, err := t.Db.Exec(`UPDATE tutor SET
	name = $1, last_name = $2, 
	email = $3, updated_at = $4, 
	WHERE tutor_id = $5
	`, tutor.Id, tutor.Name,
		tutor.LastName, tutor.Email,
		tutor.UpdatedAt, tutor.Id,
	)

	return err
}

func (t *tutorRepository) DeleteTutor(id string) error {
	_, err := t.Db.Exec(`DELETE FROM tutor WHERE tutor_id = $1`, id)
	return err
}
