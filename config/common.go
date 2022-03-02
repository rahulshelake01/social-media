package config

import "os"

type AppConfig struct {
	ENVIRONMENT string
}

var appConfig AppConfig

func initAppConfig() {
	appConfig = AppConfig{
		ENVIRONMENT: os.Getenv("ENV"),
	}
}

func GetAppConfig() AppConfig {
	return appConfig
}
