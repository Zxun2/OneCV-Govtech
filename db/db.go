package db

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect connects to the database
func Connect(dbSource string, logLevel logger.LogLevel)(*gorm.DB, error) {
	var err error
	dbDriver := mysql.Open(dbSource)
	gormCfg := newGormConfig(logLevel)
	db, err := gorm.Open(dbDriver, gormCfg)
	if err != nil {
		return nil, err
	}
	return db, nil
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
