package main

import (
	"Zxun2/OneCV-Govtech/api"
	"Zxun2/OneCV-Govtech/db"
	"Zxun2/OneCV-Govtech/models"
	"Zxun2/OneCV-Govtech/seed"
	"Zxun2/OneCV-Govtech/utils"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	switch config.Environment {
		case "development":
			config.LogLevel = logger.Info
		case "test":
			config.LogLevel = logger.Silent
		case "production":
			config.LogLevel = logger.Warn
		default:
	}

	database, err := db.Connect(config)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}
	log.Printf("Connected to database: %s", config.DatabaseURL)

	database.AutoMigrate(&models.Teacher{}, &models.Student{})
	log.Printf("Successfully migrated models")

	seed.SeedData()

	runServer(config, database)

}

func runServer(config utils.Config, store *gorm.DB) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("Cannot start server")
	}
}
