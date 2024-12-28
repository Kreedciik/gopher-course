package repository

import (
	"database/sql"
	"lesson22/model"

	"github.com/google/uuid"
)

type TutorRepository struct {
	Db *sql.DB
}

func CreateTutorRepository(db *sql.DB) TutorRepository {
	return TutorRepository{Db: db}
}

func (t *TutorRepository) CreateTutor(tutor model.Tutor) error {
	_, err := t.Db.Exec(`INSERT INTO tutors (
		tutor_id, name, last_name, 
		email
		) VALUES 
		($1, $2, $3, $4)`,
		uuid.New(), tutor.Name,
		tutor.LastName, tutor.Email,
	)
	return err
}

func (t *TutorRepository) GetTutor(id string) (model.Tutor, error) {
	var tutor model.Tutor
	row := t.Db.QueryRow(`SELECT tutor_id, name, last_name, email
		FROM tutors
		WHERE tutor_id = $1
	`, id)

	err := row.Scan(
		&tutor.Id, &tutor.Name,
		&tutor.LastName, &tutor.Email,
	)

	return tutor, err
}

func (t *TutorRepository) GetAllTutors() ([]model.Tutor, error) {
	var tutors []model.Tutor

	rows, err := t.Db.Query(
		`SELECT tutor_id, name, last_name, email FROM tutors`)

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

func (t *TutorRepository) UpdateTutor(tutor model.Tutor) error {
	_, err := t.Db.Exec(`UPDATE tutors SET
	name = $1, last_name = $2, 
	email = $3, updated_at = $4, 
	WHERE tutor_id = $5
	`, tutor.Id, tutor.Name,
		tutor.LastName, tutor.Email,
		tutor.UpdatedAt, tutor.Id,
	)

	return err
}

func (t *TutorRepository) DeleteTutor(id string) error {
	_, err := t.Db.Exec(`DELETE FROM tutors WHERE tutor_id = $1`, id)
	return err
}
