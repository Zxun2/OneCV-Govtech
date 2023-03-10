package tests

import (
	"Zxun2/OneCV-Govtech/models"
	"Zxun2/OneCV-Govtech/utils"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)	

var testDb *gorm.DB

func TestMain(m *testing.M) {
	var err error
	config, err := utils.LoadConfig("../")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	sqlDB, _ := sql.Open(config.DatabaseDriver,config.TestDatabaseURL)
	testDb, err = gorm.Open(mysql.New(mysql.Config{
  	Conn: sqlDB,
	}), &gorm.Config{})


	err = testDb.AutoMigrate(&models.Teacher{}, &models.Student{})
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	os.Exit(m.Run())
}


