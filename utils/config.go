package utils

import (
	"github.com/spf13/viper"
	"gorm.io/gorm/logger"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Environment					string        `mapstructure:"ENVIRONMENT"`
	DatabaseURL 			 	string        `mapstructure:"DATABASE_URL"`
	TestDatabaseURL 		string        `mapstructure:"TEST_DATABASE_URL"`
	DatabaseDriver 			string        `mapstructure:"DATABASE_DRIVER"`
	HTTPServerAddress   string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	LogLevel     				logger.LogLevel
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}