package postgres

import (
	"hw20/model"

	"gorm.io/gorm"
)

func DoAutoMigration(orm *gorm.DB) {
	orm.AutoMigrate(&model.Student{})
	orm.AutoMigrate(&model.Tutor{})
	orm.AutoMigrate(&model.Course{})
	orm.AutoMigrate(&model.Enroll{})
	orm.AutoMigrate(&model.Group{})
}
