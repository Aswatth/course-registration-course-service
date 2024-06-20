package services

import (
	"course-registration-system/course-service/models"
	"errors"
)

type CourseCrudService struct {
	sqlDatabase MySqlDatabase
}

func (obj *CourseCrudService) Init(db MySqlDatabase) {
	obj.sqlDatabase = db
	obj.sqlDatabase.db.AutoMigrate(&models.Course{})
}

func (obj *CourseCrudService) CreateCourse(course models.Course) error {
	result := obj.sqlDatabase.db.Create(&course)

	return result.Error
}

func (obj *CourseCrudService) FetchCourse(course_id int) (models.Course, error) {
	var course models.Course

	result := obj.sqlDatabase.db.First(&course, course_id)

	return course, result.Error
}

func (obj *CourseCrudService) UpdateCourse(course models.Course) error {
	result := obj.sqlDatabase.db.Model(&models.Course{}).Where("course_id = ?", course.Course_id).Updates(course)

	if result.RowsAffected == 0 {
		return errors.New("record not found / no updates")
	}

	return result.Error
}

func (obj *CourseCrudService) DeleteCourse(course_id int) error {
	result := obj.sqlDatabase.db.Delete(&models.Course{}, course_id)

	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return result.Error
}
