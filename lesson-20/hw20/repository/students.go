package repository

import (
	"hw20/model"

	"gorm.io/gorm"
)

type studentRepository struct {
	orm *gorm.DB
}

func CreateStudentRepository(orm *gorm.DB) studentRepository {
	return studentRepository{orm}
}

func (s *studentRepository) CreateStudent(student model.Student) error {
	err := s.orm.Create(&student).Error
	return err
}

func (s *studentRepository) GetStudent(id string) (model.Student, error) {
	var student model.Student
	err := s.orm.First(&student, id).Error
	return student, err
}

func (s *studentRepository) GetAllStudents() ([]model.Student, error) {
	var students []model.Student
	err := s.orm.Find(&students).Error
	return students, err
}

func (s *studentRepository) UpdateStudent(student model.Student) error {
	err := s.orm.Save(&student).Error
	return err
}

func (s *studentRepository) DeleteStudent(id string) error {
	err := s.orm.Delete(model.Student{}, id).Error
	return err
}

func (s *studentRepository) GetOldestStudent() (model.Student, error) {
	var student model.Student
	err := s.orm.Limit(1).Order("age DESC").Find(&student).Error
	return student, err
}

func (s *studentRepository) GetYoungestStudent() (model.Student, error) {
	var student model.Student
	err := s.orm.Limit(1).Order("age ASC").Find(&student).Error
	return student, err
}

func (s *studentRepository) GetBestStudentByGroup(groupId string) (model.Student, error) {
	var student model.Student
	err := s.orm.Table("groups").Select(`groups.student_id, 
	groups.name, groups.lastname, 
	groups.phone, groups.age, 
	groups.grade, groups.gender`).Joins(`
		JOIN courses ON groups.course_id = courses.course_id
	`).Joins(`
		JOIN enrollments ON courses.course_id = enrollments.course_id
	`).Joins(`
		JOIN students ON enrollments.student_id = students.student_id
	`).Where("group_id = ", groupId).Order("grade DESC").Limit(1).Find(&student).Error

	return student, err
}

func (s *studentRepository) GetBestStudentByCourse(courseID string) (model.Student, error) {
	var student model.Student
	err := s.orm.Table(`students
	`).Select(`students.student_id, students.name, 
	students.lastname, students.phone, 
	students.age, students.grade, students.gender
	`).Joins(`
		JOIN enrollments ON students.student_id = enrollments.student_id
	`).Joins(`
		JOIN courses ON enrollments.course_id = courses.course_id
	`).Where("courses.courses_id = ", courseID).Order("grade DESC").Limit(1).Find(&student).Error
	return student, err
}

func (s *studentRepository) GetStudentsByCourse(courseID string) (model.Student, error) {
	var student model.Student
	err := s.orm.Raw(`
			SELECT s.student_id, s.name, s.lastname, s.phone, s.age, s.grade, s.gender 
			FROM students s
			JOIN enrollments e USING(student_id)
			JOIN courses c USING(course_id)
			WHERE c.course_id = ?
	`, courseID).Scan(&student).Error
	return student, err
}

func (s *studentRepository) GetStudentsByGender(courseID string, gender string) (model.Student, error) {
	var student model.Student
	err := s.orm.Raw(`
			SELECT s.student_id, s.name, s.lastname, s.phone, s.age, s.grade, s.gender 
			FROM students s
			JOIN enrollments e USING(student_id)
			JOIN courses c USING(course_id)
			WHERE c.course_id = ? AND gender = ?
	`, courseID, gender).Scan(&student).Error
	return student, err
}

func (s *studentRepository) GetStudentsByGroup(groupID string, gender string) (model.Student, error) {
	var student model.Student
	err := s.orm.Raw(`
			SELECT s.student_id, s.name, s.lastname, s.phone, s.age, s.grade, s.gender 
			FROM students s
			JOIN enrollments e USING(student_id)
			JOIN courses c USING(course_id)
			JOIN groups g USING(course_id)
			WHERE g.group_id = ?
			AND s.gender = ?
	`, groupID, gender).Scan(&student).Error
	return student, err
}

func (s *studentRepository) GetStudentsByTutor(tutorID string, gender string) (model.Student, error) {
	var student model.Student
	err := s.orm.Raw(`
			SELECT s.student_id, s.name, s.lastname, s.phone, s.age, s.grade, s.gender 
			FROM students s
			JOIN enrollments e USING(student_id)
			JOIN courses c USING(course_id)
			JOIN tutors t USING(tutor_id)
			WHERE t.tutor_id = ?
			AND s.gender = ?
	`, tutorID, gender).Scan(&student).Error
	return student, err
}
