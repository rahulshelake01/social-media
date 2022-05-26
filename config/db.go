package config

import (
	"os"
	"strconv"
)

type DatabaseConfig struct {
	Host               string
	Port               int
	User               string
	Password           string
	DBName             string
	MaxPoolSize        int
	MaxIdleConnections int
}

var dbConfig *DatabaseConfig

func initDBConfig() {

	MaxPoolSize, _ := strconv.Atoi(os.Getenv("MYSQL_MAX_POOL_SIZE"))
	MaxIdleConnections, _ := strconv.Atoi(os.Getenv("MYSQL_MAX_IDLE_CONNECTIONS"))
	Port, _ := strconv.Atoi(os.Getenv("MYSQL_PORT"))

	dbConfig = &DatabaseConfig{
		Host:               os.Getenv("MYSQL_HOST"),
		Port:               Port,
		User:               os.Getenv("MYSQL_USER"),
		Password:           os.Getenv("MYSQL_PASSWORD"),
		DBName:             os.Getenv("MYSQL_DATABASE"),
		MaxPoolSize:        MaxPoolSize,
		MaxIdleConnections: MaxIdleConnections,
	}
}

func GetDBConfig() *DatabaseConfig {
	return dbConfig
}
