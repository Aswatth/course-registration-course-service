package controllers

import (
	"course-registration-system/course-service/models"
	"course-registration-system/course-service/services"

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
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
	} else {
		//Store to DB
		err := obj.course_crud_service.CreateCourse(course)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		} else {
			context.Status(http.StatusOK)
		}
	}

}

func (obj *CourseCrudController) FetchCourse(context *gin.Context) {
	course_id, err := strconv.ParseInt(context.Param("course_id"), 0, 0)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
	} else {
		//Fetch from DB
		fetched_course, err := obj.course_crud_service.FetchCourse(int(course_id))

		if err != nil {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": err.Error()})
		} else {
			context.JSON(http.StatusOK, fetched_course)
		}
	}
}

func (obj *CourseCrudController) UpdateCourse(context *gin.Context) {

	//Fetch course
	course_id, err := strconv.ParseInt(context.Param("course_id"), 0, 0)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		var course models.Course

		//Check if given JSON is valid
		if err := context.ShouldBindJSON(&course); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			course.Course_id = int(course_id)

			err := obj.course_crud_service.UpdateCourse(course)

			if err != nil {
				context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": err.Error()})
			} else {
				context.Status(http.StatusOK)
			}
		}
	}
}

func (obj *CourseCrudController) DeleteCourse(context *gin.Context) {
	course_id, err := strconv.ParseInt(context.Param("course_id"), 0, 0)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
	} else {
		err := obj.course_crud_service.DeleteCourse(int(course_id))

		if err != nil {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": err.Error()})
		} else {
			context.Status(http.StatusOK)
		}
	}
}

func (obj *CourseCrudController) RegisterRoutes(rg *gin.RouterGroup) {
	course_routes := rg.Group("/courses")
	course_routes.POST("", obj.CreateCourse)
	course_routes.GET("/:course_id", obj.FetchCourse)
	course_routes.PUT("/:course_id", obj.UpdateCourse)
	course_routes.DELETE("/:course_id", obj.DeleteCourse)
}
