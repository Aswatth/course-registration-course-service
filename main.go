package main

import (
	"context"
	"course-registration-system/course-service/controllers"
	"course-registration-system/course-service/models"
	"course-registration-system/course-service/services"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Database struct {
	DB_TYPE              string
	DB_NAME              string
	DB_CONNECTION_STRING string
}

type Config struct {
	PORT int
	DB   []Database
}

func (config *Config) LoadConfig(file_name string) {

	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal("Error occured while reading config file.", err)
	}

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Error occured while reading config file.", err)
	}
}

func ConnectToDB(config Config) (services.MongoDatabase, services.MySqlDatabase) {
	sql_database := new(services.MySqlDatabase)
	mongo_database := new(services.MongoDatabase)

	for _, database := range config.DB {
		if database.DB_TYPE == "MONGO" {
			mongo_database.Connect(context.Background(), database.DB_CONNECTION_STRING)
			defer mongo_database.Disconnect(context.Background())
			mongo_database.Ping(context.Background())
			mongo_database.SetDatabase(database.DB_NAME)
			fmt.Println("Connected to Mongo DB")
		} else if database.DB_TYPE == "MYSQL" {
			sql_database.Connect(database.DB_CONNECTION_STRING)
			fmt.Println("Connected to MySQL DB")
		}
	}

	return *mongo_database, *sql_database
}

func InitializeCourseCrud(mongo_db services.MongoDatabase, mysql_db services.MySqlDatabase) *controllers.CourseCrudController {
	course_service := new(services.CourseCrudService)
	course_service.Init(mysql_db)

	course_controller := new(controllers.CourseCrudController)
	course_controller.Init(*course_service)

	return course_controller
}

func main() {

	//Load config
	config := new(Config)
	config.LoadConfig("config.json")

	mongo_database, sql_database := ConnectToDB(*config)

	course_crud_controller := InitializeCourseCrud(mongo_database, sql_database)

	server := gin.Default()

	base_path := server.Group("")
	course_crud_controller.RegisterRoutes(base_path)

	course := new(models.Course)
	course.CreateCourse(101, "Pqr", 4, "intro to pqr", "CS")
	// fmt.Println(course)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.Run(":" + fmt.Sprint(config.PORT)) // listen and serve on 0.0.0.0:8080
}
