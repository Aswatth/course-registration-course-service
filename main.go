package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Config struct {
	PORT                  int
	DB_CONNECNTION_STRING string
	DB_NAME               string
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

	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.Run(":" + fmt.Sprint(config.PORT)) // listen and serve on 0.0.0.0:8080
}
