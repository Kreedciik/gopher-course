package repository

import (
	"fmt"
	"hw20/model"

	"github.com/savsgio/gotils/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type courseRepository struct {
	orm *gorm.DB
}

func CreateCourseRepository(orm *gorm.DB) courseRepository {
	return courseRepository{orm}
}

func (c *courseRepository) CreateCourse(course model.Course) error {
	result := c.orm.Create(&course)
	return result.Error
}

func (c *courseRepository) GetCourse(id string) (model.Course, error) {
	var course model.Course
	tx := c.orm.First(&course, id)
	return course, tx.Error
}

func (c *courseRepository) GetCourses() ([]model.Course, error) {
	var courses []model.Course
	tx := c.orm.Find(&courses)
	return courses, tx.Error
}

func (c *courseRepository) UpdateCourse(course model.Course) error {
	tx := c.orm.Save(&course)
	return tx.Error
}

func (c *courseRepository) DeleteCourse(id string) error {
	tx := c.orm.Delete(&model.Course{}, id)
	return tx.Error
}

func (c *courseRepository) GetBiggestCourse() (model.Course, error) {
	var course model.Course
	tx := c.orm.Limit(1).Order("student_number DESC").Find(&course)
	return course, tx.Error
}

func (c *courseRepository) EnrollToCourse(studentID string, courseID string) error {

	tx := c.orm.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var id string
	err := tx.Select("student_id").First(&id).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if len(id) > 0 {
		tx.Rollback()
		err := fmt.Errorf("the student with id %s already enrolled", studentID)
		return err
	}

	enroll := model.Enroll{CourseID: courseID, StudentID: id, Id: uuid.V4()}
	err = tx.Create(&enroll).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Model(&model.Course{}).Where("course_id = ?", courseID).Update("student_number", gorm.Expr("student_number + ?", 1)).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (c *courseRepository) ExcludeFromCourse(studentID string) error {
	tx := c.orm.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var course_id string
	err := tx.Model(model.Enroll{}).Clauses(clause.Returning{Columns: []clause.Column{{Name: "course_id"}}}).Where("student_id = ?", studentID).Delete(&course_id).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Model(model.Course{}).Where("course_id = ?", course_id).Update("student_number", gorm.Expr("student_number - ?", 1)).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
