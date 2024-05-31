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
