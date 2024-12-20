package repository

import (
	"hw20/model"

	"gorm.io/gorm"
)

type tutorRepository struct {
	orm *gorm.DB
}

func CreateTutorRepository(orm *gorm.DB) tutorRepository {
	return tutorRepository{orm}
}

func (t *tutorRepository) CreateTutor(tutor model.Tutor) error {
	err := t.orm.Create(&tutor).Error
	return err
}

func (t *tutorRepository) GetTutor(id string) (model.Tutor, error) {
	var tutor model.Tutor
	err := t.orm.First(&tutor, id).Error
	return tutor, err
}

func (t *tutorRepository) GetAllTutors() ([]model.Tutor, error) {
	var tutors []model.Tutor
	err := t.orm.Find(&tutors).Error
	return tutors, err
}

func (t *tutorRepository) UpdateTutor(tutor model.Tutor) error {
	err := t.orm.Save(&tutor).Error
	return err
}

func (t *tutorRepository) DeleteTutor(id string) error {
	err := t.orm.Delete(&model.Tutor{}, id).Error
	return err
}
