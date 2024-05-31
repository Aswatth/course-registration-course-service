package models

type Course struct {
	course_id          int
	course_name        string
	credits            int
	course_description string
	department         string
}

func (obj *Course) CreateCourse(course_id int, course_name string, credits int, course_description string, department string) {
	obj.course_id = course_id
	obj.course_name = course_name
	obj.credits = credits
	obj.course_description = course_description
	obj.department = department
}
