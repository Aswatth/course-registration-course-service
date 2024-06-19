package main

import (
	"course-registration-system/course-service/controllers"
	"course-registration-system/course-service/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sql_database := new(services.MySqlDatabase)
	sql_database.Connect(os.Getenv("MYSQL_CONNECTION_STRING"))

	course_service := new(services.CourseCrudService)
	course_service.Init(*sql_database)

	course_controller := new(controllers.CourseCrudController)
	course_controller.Init(*course_service)

	server := gin.Default()

	base_path := server.Group("")
	course_controller.RegisterRoutes(base_path)

	server.Run(":" + os.Getenv("PORT")) // listen and serve on 0.0.0.0:8080
}
