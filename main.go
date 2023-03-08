package main

import (
	"Zxun2/OneCV-Govtech/controllers"
	"Zxun2/OneCV-Govtech/db"
	"Zxun2/OneCV-Govtech/models"
	"Zxun2/OneCV-Govtech/seed"
	"Zxun2/OneCV-Govtech/utils"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	database, err := db.Connect(config.DatabaseURL)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}
	log.Printf("Connected to database: %s", config.DatabaseURL)

	database.AutoMigrate(&models.Teacher{}, &models.Student{})
	log.Printf("Successfully migrated models")

	seed.SeedData()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	api := router.Group("/api")

	api.POST("/register", controllers.Register)
	api.GET("/commonstudents", controllers.GetCommonStudents)
	api.POST("/suspend", controllers.Suspend)
	api.POST("/retrievefornotifications", controllers.RetrieveNotifications)

	router.Run(config.HTTPServerAddress)
}
