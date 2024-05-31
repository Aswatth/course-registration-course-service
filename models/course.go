package models

type Course struct {
	Course_id          int `gorm:"primaryKey"`
	Course_name        string
	Credits            int
	Course_description string
	Department         string
}

func (obj *Course) CreateCourse(course_id int, course_name string, credits int, course_description string, department string) {
	obj.Course_id = course_id
	obj.Course_name = course_name
	obj.Credits = credits
	obj.Course_description = course_description
	obj.Department = department
}
