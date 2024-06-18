package services

import (
	"course-registration-system/course-service/models"
	"fmt"
	"log"
)

type CourseCrudService struct {
	sqlDatabase MySqlDatabase
}

func (obj *CourseCrudService) Init(db MySqlDatabase) {
	obj.sqlDatabase = db
	obj.sqlDatabase.db.AutoMigrate(&models.Course{})
}

func (obj *CourseCrudService) CreateCourse(course models.Course) {
	result := obj.sqlDatabase.db.Create(&course)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	if result.RowsAffected > 0 {
		fmt.Println("New course created successfully")
	}
}

func (obj *CourseCrudService) FetchCourse(course_id int) models.Course {
	var course models.Course

	obj.sqlDatabase.db.First(&course, course_id)

	return course
}

func (obj *CourseCrudService) UpdateCourse(course models.Course) {
	obj.sqlDatabase.db.Model(&models.Course{}).Where("course_id = ?", course.Course_id).Updates(course)
}

func (obj *CourseCrudService) DeleteCourse(course_id int) {
	var course models.Course

	obj.sqlDatabase.db.Delete(&course, course_id)
}
