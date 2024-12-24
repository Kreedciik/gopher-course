package repository

import (
	"database/sql"
	"lesson22/model"
)

type studentRepository struct {
	db *sql.DB
}

func CreateStudentRepository(db *sql.DB) studentRepository {
	return studentRepository{db}
}

func (s *studentRepository) CreateStudent(student model.Student) error {
	_, err := s.db.Exec(`INSERT INTO student VALUES (
	name, lastname, phone, age, grade, gender 
	)`,
		student.Name, student.LastName,
		student.Phone, student.Age,
		student.Grade, student.Gender,
	)
	return err
}

func (s *studentRepository) GetStudent(id string) (model.Student, error) {
	var student model.Student
	row := s.db.QueryRow(`SELECT student_id, name, lastname, phone, age, grade, gender
		FROM student
		WHERE student_id = $1
	`, id)

	err := row.Scan(
		&student.StudentID, &student.Name,
		&student.LastName, &student.Phone,
		&student.Age, &student.Grade,
		&student.Gender)

	return student, err
}

func (s *studentRepository) GetAllStudents() ([]model.Student, error) {
	var students []model.Student

	rows, err := s.db.Query(`SELECT student_id, name, lastname, phone, age, grade, gender FROM student`)

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

func (s *studentRepository) UpdateStudent(student model.Student) error {
	_, err := s.db.Exec(`UPDATE student SET
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

func (s *studentRepository) DeleteStudent(id string) error {
	_, err := s.db.Exec(`DELETE FROM student WHERE student_id = $1`, id)
	return err
}

func (s *studentRepository) GetOldestStudent() (model.Student, error) {
	var student model.Student
	row := s.db.QueryRow(`SELECT student_id, name, lastname, phone, age, grade, gender
		FROM student
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

func (s *studentRepository) GetYoungestStudent() (model.Student, error) {
	var student model.Student
	row := s.db.QueryRow(`SELECT student_id, name, lastname, phone, age, grade, gender
		FROM student
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

func (s *studentRepository) GetBestStudentByGroup(groupId string) (model.Student, error) {
	var student model.Student
	row := s.db.QueryRow(`
			SELECT s.student_id, s.name, s.lastname, s.phone, s.age, s.grade, s.gender 
			FROM student_group g
			JOIN course c USING(course_id)
			JOIN enrollment e USING(course_id)
			JOIN student s USING (student_id)
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

func (s *studentRepository) GetBestStudentByCourse(courseID string) (model.Student, error) {
	var student model.Student
	row := s.db.QueryRow(`
			SELECT s.student_id, s.name, s.lastname, s.phone, s.age, s.grade, s.gender 
			FROM student s
			JOIN enrollment e USING(student_id)
			JOIN course c USING(course_id)
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

func (s *studentRepository) GetStudentsByCourse(courseID string) (model.Student, error) {
	var student model.Student
	row := s.db.QueryRow(`
			SELECT s.student_id, s.name, s.lastname, s.phone, s.age, s.grade, s.gender 
			FROM student s
			JOIN enrollment e USING(student_id)
			JOIN course c USING(course_id)
			WHERE c.course_id = $1
	`, courseID)

	err := row.Scan(
		&student.StudentID, &student.Name,
		&student.LastName, &student.Phone,
		&student.Age, &student.Grade,
		&student.Gender)

	return student, err
}

func (s *studentRepository) GetStudentsByGender(courseID string, gender string) (model.Student, error) {
	var student model.Student
	row := s.db.QueryRow(`
			SELECT s.student_id, s.name, s.lastname, s.phone, s.age, s.grade, s.gender 
			FROM student s
			JOIN enrollment e USING(student_id)
			JOIN course c USING(course_id)
			WHERE c.course_id = $1 AND gender = $2
	`, courseID, gender)

	err := row.Scan(
		&student.StudentID, &student.Name,
		&student.LastName, &student.Phone,
		&student.Age, &student.Grade,
		&student.Gender)

	return student, err
}

func (s *studentRepository) GetStudentsByGroup(groupID string, gender string) (model.Student, error) {
	var student model.Student
	row := s.db.QueryRow(`
			SELECT s.student_id, s.name, s.lastname, s.phone, s.age, s.grade, s.gender 
			FROM student s
			JOIN enrollment e USING(student_id)
			JOIN course c USING(course_id)
			JOIN student_group g USING(course_id)
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

func (s *studentRepository) GetStudentsByTutor(tutorID string, gender string) (model.Student, error) {
	var student model.Student
	row := s.db.QueryRow(`
			SELECT s.student_id, s.name, s.lastname, s.phone, s.age, s.grade, s.gender 
			FROM student s
			JOIN enrollment e USING(student_id)
			JOIN course c USING(course_id)
			JOIN tutor t USING(tutor_id)
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
