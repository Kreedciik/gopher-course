package repository

import (
	"database/sql"
	"fmt"
	"lesson22/model"
)

type StudentRepository struct {
	db *sql.DB
}

func CreateStudentRepository(db *sql.DB) StudentRepository {
	return StudentRepository{db}
}

func (s *StudentRepository) CreateStudent(student model.CreateStudentRequest) error {
	_, err := s.db.Exec(`INSERT INTO students VALUES (
	name, lastname, phone, age, grade, gender 
	), $1, $2, $3, $4, $5, $6`,
		student.Name, student.LastName,
		student.Phone, student.Age,
		student.Grade, student.Gender,
	)
	return err
}

func (s *StudentRepository) GetStudent(id string) (model.Student, error) {
	var student model.Student
	row := s.db.QueryRow(`SELECT student_id, name, lastname, phone, age, grade, gender
		FROM students
		WHERE student_id = $1
	`, id)

	err := row.Scan(
		&student.StudentID, &student.Name,
		&student.LastName, &student.Phone,
		&student.Age, &student.Grade,
		&student.Gender)

	return student, err
}

func (s *StudentRepository) GetAllStudents() ([]model.Student, error) {
	var students []model.Student

	rows, err := s.db.Query(`SELECT student_id, name, lastname, phone, age, grade, gender FROM students`)

	for rows.Next() {
		var student model.Student
		err = rows.Scan(
			&student.StudentID, &student.Name,
			&student.LastName, &student.Phone,
			&student.Age, &student.Grade,
			&student.Gender)
		students = append(students, student)
	}
	return students, err
}

func (s *StudentRepository) UpdateStudent(student model.Student) error {
	_, err := s.db.Exec(`UPDATE students SET
		name = $1, lastname = $2, 
		phone = $3, age = $4, 
		grade = $5, gender = $6
		updated_at = $7
		WHERE student_id = $8
	`, student.Name, student.LastName,
		student.Phone, student.Age,
		student.Grade, student.Gender,
		student.UpdatedAt, student.StudentID,
	)

	return err
}

func (s *StudentRepository) DeleteStudent(id string) error {
	r, err := s.db.Exec(`DELETE FROM students WHERE student_id = $1`, id)
	if err != nil {
		return err
	}

	n, err := r.RowsAffected()
	if n == 0 || err != nil {
		return fmt.Errorf("not listed")
	}
	return nil
}

func (s *StudentRepository) GetOldestStudent() (model.Student, error) {
	var student model.Student
	row := s.db.QueryRow(`SELECT student_id, name, lastname, phone, age, grade, gender
		FROM students
		ORDER BY age DESC
		LIMIT 1;
	`)

	err := row.Scan(
		&student.StudentID, &student.Name,
		&student.LastName, &student.Phone,
		&student.Age, &student.Grade,
		&student.Gender)

	return student, err
}

func (s *StudentRepository) GetYoungestStudent() (model.Student, error) {
	var student model.Student
	row := s.db.QueryRow(`SELECT student_id, name, lastname, phone, age, grade, gender
		FROM students
		ORDER BY age
		LIMIT 1;
	`)

	err := row.Scan(
		&student.StudentID, &student.Name,
		&student.LastName, &student.Phone,
		&student.Age, &student.Grade,
		&student.Gender)

	return student, err
}

func (s *StudentRepository) GetBestStudentByGroup(groupId string) (model.Student, error) {
	var student model.Student
	row := s.db.QueryRow(`
			SELECT s.student_id, s.name, s.lastname, s.phone, s.age, s.grade, s.gender 
			FROM groups g
			JOIN courses c USING(course_id)
			JOIN enrollments e USING(course_id)
			JOIN students s USING (student_id)
			WHERE g.group_id = $1
			ORDER BY grade DESC
			LIMIT 1;
	`, groupId)

	err := row.Scan(
		&student.StudentID, &student.Name,
		&student.LastName, &student.Phone,
		&student.Age, &student.Grade,
		&student.Gender)

	return student, err
}

func (s *StudentRepository) GetBestStudentByCourse(courseID string) (model.Student, error) {
	var student model.Student
	row := s.db.QueryRow(`
			SELECT s.student_id, s.name, s.lastname, s.phone, s.age, s.grade, s.gender 
			FROM students s
			JOIN enrollments e USING(student_id)
			JOIN courses c USING(course_id)
			WHERE c.course_id = $1
			ORDER BY grade DESC
			LIMIT 1;
	`, courseID)

	err := row.Scan(
		&student.StudentID, &student.Name,
		&student.LastName, &student.Phone,
		&student.Age, &student.Grade,
		&student.Gender)

	return student, err
}

func (s *StudentRepository) GetStudentsByCourse(courseID string) (model.Student, error) {
	var student model.Student
	row := s.db.QueryRow(`
			SELECT s.student_id, s.name, s.lastname, s.phone, s.age, s.grade, s.gender 
			FROM students s
			JOIN enrollments e USING(student_id)
			JOIN courses c USING(course_id)
			WHERE c.course_id = $1
	`, courseID)

	err := row.Scan(
		&student.StudentID, &student.Name,
		&student.LastName, &student.Phone,
		&student.Age, &student.Grade,
		&student.Gender)

	return student, err
}

func (s *StudentRepository) GetStudentsByGender(courseID string, gender string) (model.Student, error) {
	var student model.Student
	row := s.db.QueryRow(`
			SELECT s.student_id, s.name, s.lastname, s.phone, s.age, s.grade, s.gender 
			FROM students s
			JOIN enrollments e USING(student_id)
			JOIN courses c USING(course_id)
			WHERE c.course_id = $1 AND gender = $2
	`, courseID, gender)

	err := row.Scan(
		&student.StudentID, &student.Name,
		&student.LastName, &student.Phone,
		&student.Age, &student.Grade,
		&student.Gender)

	return student, err
}

func (s *StudentRepository) GetStudentsByGroup(groupID string, gender string) (model.Student, error) {
	var student model.Student
	row := s.db.QueryRow(`
			SELECT s.student_id, s.name, s.lastname, s.phone, s.age, s.grade, s.gender 
			FROM students s
			JOIN enrollments e USING(student_id)
			JOIN courses c USING(course_id)
			JOIN groups g USING(course_id)
			WHERE g.group_id = $1
			AND s.gender = $2
	`, groupID, gender)

	err := row.Scan(
		&student.StudentID, &student.Name,
		&student.LastName, &student.Phone,
		&student.Age, &student.Grade,
		&student.Gender)

	return student, err
}

func (s *StudentRepository) GetStudentsByTutor(tutorID string, gender string) (model.Student, error) {
	var student model.Student
	row := s.db.QueryRow(`
			SELECT s.student_id, s.name, s.lastname, s.phone, s.age, s.grade, s.gender 
			FROM students s
			JOIN enrollments e USING(student_id)
			JOIN courses c USING(course_id)
			JOIN tutors t USING(tutor_id)
			WHERE t.tutor_id = $1
			AND s.gender = $2
	`, tutorID, gender)

	err := row.Scan(
		&student.StudentID, &student.Name,
		&student.LastName, &student.Phone,
		&student.Age, &student.Grade,
		&student.Gender)

	return student, err
}
