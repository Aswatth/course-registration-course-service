package controllers

import (
	"course-registration-system/course-service/models"
	"course-registration-system/course-service/services"

	// "fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CourseCrudController struct {
	course_crud_service services.CourseCrudService
}

func (obj *CourseCrudController) Init(course_service services.CourseCrudService) {
	obj.course_crud_service = course_service
}

func (obj *CourseCrudController) CreateCourse(context *gin.Context) {
	var course models.Course

	//Check if given JSON is valid
	if err := context.ShouldBindJSON(&course); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	// fmt.Print("[OUTPUT]:\t")
	// fmt.Println(course)

	//Store to DB
	obj.course_crud_service.CreateCourse(course)

	context.JSON(http.StatusOK, gin.H{"message": "Successfully created a course"})
}

func (obj *CourseCrudController) FetchCourse(context *gin.Context) {
	course_id, _ := strconv.ParseInt(context.Query("course_id"), 0, 0)

	//Fetch from DB
	fetched_course := obj.course_crud_service.FetchCourse(int(course_id))

	context.JSON(http.StatusOK, fetched_course)
}

func (obj *CourseCrudController) RegisterRoutes(rg *gin.RouterGroup) {
	course_routes := rg.Group("/courses")
	course_routes.POST("/create", obj.CreateCourse)
	course_routes.GET("/fetch", obj.FetchCourse)
}
