package repository

import (
	"hw20/model"

	"gorm.io/gorm"
)

type groupRepository struct {
	orm *gorm.DB
}

func CreateGroupRepository(db *gorm.DB) groupRepository {
	return groupRepository{db}
}

func (g *groupRepository) CreateGroup(group model.Group) error {
	err := g.orm.Create(&group).Error
	return err
}

func (g *groupRepository) GetGroup(id string) (model.Group, error) {
	var group model.Group
	err := g.orm.First(&group, id).Error
	return group, err
}

func (g *groupRepository) GetAllGroups() ([]model.Group, error) {
	var groups []model.Group
	err := g.orm.Find(&groups).Error
	return groups, err
}

func (g *groupRepository) UpdateGroup(group model.Group) error {
	err := g.orm.Save(&group).Error
	return err
}

func (g *groupRepository) DeleteGroup(id string) error {
	err := g.orm.Delete(&model.Group{}, id).Error
	return err
}

func (g *groupRepository) GetBiggestGroup() (model.Group, error) {
	var group model.Group
	err := g.orm.Limit(1).Find(&group).Order("student_count ASC").Error
	return group, err
}
