package repository

import (
	"database/sql"
	"fmt"
	"hw19/model"

	"github.com/savsgio/gotils/uuid"
)

type courseRepository struct {
	db *sql.DB
}

func CreateCourseRepository(db *sql.DB) courseRepository {
	return courseRepository{db}
}

func (c *courseRepository) CreateCourse(course model.Course) error {
	_, err := c.db.Exec(`INSERT INTO course (
		course_id, name, student_number, 
		tutor_id, started_at
		) VALUES 
		 ($1, $2, $3, $4, $5)`,
		course.Id, course.Name, course.StudentNumber,
		course.TutorID, course.StartedAt,
	)
	return err
}

func (c *courseRepository) GetCourse(id string) (model.Course, error) {
	var course model.Course
	row := c.db.QueryRow(`SELECT course_id, name,
		student_number, tutor_id, started_at
		FROM course
		WHERE course_id = $1
	`, id)

	err := row.Scan(
		&course.Id, &course.Name,
		&course.StudentNumber, &course.TutorID,
		&course.StartedAt)

	return course, err
}

func (c *courseRepository) GetCourses() ([]model.Course, error) {
	var courses []model.Course

	rows, err := c.db.Query(
		`SELECT course_id, name,
		student_number, tutor_id, started_at
		FROM course`)

	for rows.Next() {
		var course model.Course
		err = rows.Scan(
			&course.Id, &course.Name,
			&course.StudentNumber, &course.TutorID,
			&course.StartedAt)
		courses = append(courses, course)
	}
	return courses, err
}

func (c *courseRepository) UpdateCourse(course model.Course) error {
	_, err := c.db.Exec(`UPDATE course SET
	name = $1, student_number = $2, 
	tutor_id = $3, started_at = $4, 
	updated_at = $5
	WHERE course_id = $6
`, course.Name, course.StudentNumber,
		course.TutorID, course.StartedAt,
		course.UpdatedAt, course.Id,
	)

	return err
}

func (c *courseRepository) DeleteCourse(id string) error {
	_, err := c.db.Exec(`DELETE FROM course WHERE course_id = $1`, id)
	return err
}

func (c *courseRepository) GetBiggestCourse() (model.Course, error) {
	var course model.Course
	row := c.db.QueryRow(`SELECT course_id, name,
		student_number, tutor_id, started_at
		FROM course
		ORDER BY student_number
		LIMIT 1;
	`)

	err := row.Scan(
		&course.Id, &course.Name,
		&course.StudentNumber, &course.TutorID,
		&course.StartedAt)

	return course, err
}

func (c *courseRepository) EnrollToCourse(studentID string, courseID string) error {

	tx, err := c.db.Begin()

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

	var id string
	row := tx.QueryRow(`SELECT student_id FROM enrollment WHERE student_id = $1`, studentID)
	row.Scan(&id)
	if len(id) > 0 {
		err = fmt.Errorf("The student with ID %s already enrolled", studentID)
		return err
	}

	_, err = tx.Exec(`INSERT INTO enrollment (id, course_id, student_id) VALUES
			($1, $2, $3);
	`, uuid.V4(), courseID, studentID)

	if err != nil {
		return err
	}

	numberOfStudents := 0
	row = tx.QueryRow(`SELECT student_number FROM course 
	WHERE course_id = $1`, courseID)
	err = row.Scan(&numberOfStudents)
	if err != nil {
		return err
	}
	numberOfStudents++
	_, err = tx.Exec(`
	UPDATE course SET student_number = $1
	WHERE course_id = $2
	`, numberOfStudents, courseID)

	if err != nil {
		return err
	}

	return nil
}

func (c *courseRepository) ExcludeFromCourse(studentID string) error {
	tx, err := c.db.Begin()

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if err != nil {
		return err
	}

	row := tx.QueryRow(`DELETE FROM enrollment WHERE student_id = $1
				RETURNING course_id
	`, studentID)

	var course_id string
	err = row.Scan(&course_id)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`UPDATE course SET student_number = student_number - 1 WHERE course_id = $1`, course_id)

	if err != nil {
		return err
	}

	return nil
}
