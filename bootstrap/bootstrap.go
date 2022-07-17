package bootstrap

import (
	"framework/configs"
	"framework/database"

	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnvironmentVariables() configs.Config {
	config := configs.LoadConfigs()

	return config
}

func Database(config configs.Config) (*gorm.DB, error) {
	cfg := config.MySQL

	db, err := database.Connection(cfg.Username, cfg.Password, cfg.Host, cfg.Database)
	if err != nil {
		return nil, err
	}

	DB = db

	return nil, err
}
