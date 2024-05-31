package main

import (
	"context"
	"course-registration-sysyem/course-service/services"
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

func main() {

	//Load config
	config := new(Config)
	config.LoadConfig("config.json")

	sqlDatabase := new(services.MySqlDatabase)
	mongoDatabse := new(services.MongoDatabase)

	for _, database := range config.DB {
		if database.DB_TYPE == "MONGO" {
			mongoDatabse.Connect(context.Background(), database.DB_CONNECTION_STRING)
			defer mongoDatabse.Disconnect(context.Background())
			mongoDatabse.Ping(context.Background())
			mongoDatabse.SetDatabase(database.DB_NAME)
			fmt.Println("Connected to Mongo DB")
		} else if database.DB_TYPE == "MYSQL" {
			sqlDatabase.Connect(database.DB_CONNECTION_STRING)
			fmt.Println("Connected to MySQL DB")
		}
	}

	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.Run(":" + fmt.Sprint(config.PORT)) // listen and serve on 0.0.0.0:8080
}
