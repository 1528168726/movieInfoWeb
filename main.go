package main

import (
	"databaseHW/config"
	"databaseHW/models"
	"databaseHW/router"
	"log"
)

func main() {
	dbConfig, err := config.GetDbInfo("dbConfig.json")
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = models.DbInit(&dbConfig)
	if err != nil {
		log.Fatalln(err)
		return
	}
	r := router.InitRouter()
	r.Run(":8080")
}
