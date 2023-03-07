package main

import (
	"Zxun2/OneCV-Govtech/controllers"
	"Zxun2/OneCV-Govtech/models"
	"Zxun2/OneCV-Govtech/utils"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-contrib/cors"
)

// DB is the global database connection
var DB *gorm.DB

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	DB, err = gorm.Open(mysql.Open(config.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	DB.AutoMigrate(&models.Teacher{}, &models.Student{}, &models.TeacherStudent{})


	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	api := router.Group("/api")

	api.POST("/register", controllers.Placeholder)
	api.GET("/commonStudents", controllers.Placeholder)
	api.POST("/suspend", controllers.Placeholder)
	api.POST("/retrievefornotifications", controllers.Placeholder)

	router.Run(config.HTTPServerAddress)
}
