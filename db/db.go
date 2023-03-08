package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Store is the global database connection
var Store *gorm.DB

// Connect connects to the database
func Connect(databaseURL string)(*gorm.DB, error) {
	var err error
	Store, err = gorm.Open(mysql.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return Store, nil
}