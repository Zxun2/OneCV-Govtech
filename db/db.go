package db

import (
	"Zxun2/OneCV-Govtech/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Store is the global database connection
var Store *gorm.DB

// Connect connects to the database
func Connect(config utils.Config)(*gorm.DB, error) {
	var err error
	dbDriver := mysql.Open(config.DatabaseURL)
	gormCfg := newGormConfig(config.LogLevel)
	Store, err = gorm.Open(dbDriver, gormCfg)
	if err != nil {
		return nil, err
	}
	return Store, nil
}

// newGormConfig creates a gorm.Config with the specified log level.
func newGormConfig(logLevel logger.LogLevel) *gorm.Config {
	return &gorm.Config{Logger: logger.New(
		logrus.StandardLogger(),
		logger.Config{
			LogLevel: logLevel,
			Colorful: true,
		},
	)}
}
