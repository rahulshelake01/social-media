package database

import (
	"database/sql"
	"social-media/config"
	"time"

	_ "github.com/go-sql-driver/mysql"

	log "github.com/sirupsen/logrus"
)

type DBConnectionInterface interface {
	DBConnect() *sql.DB
}

type Connection struct {
	DBConfig *config.DatabaseConfig
}

func DBConnection(DBConfig *config.DatabaseConfig) DBConnectionInterface {
	return Connection{DBConfig}
}

func (db Connection) DBConnect() *sql.DB {

	dbConfigs := db.DBConfig

	dbConn, errConn := sql.Open("mysql", dbConfigs.User+":"+dbConfigs.Password+"@tcp("+dbConfigs.Host+")/"+dbConfigs.DBName)

	if errConn != nil {
		log.Error("Failed to connect to db : ", errConn)
		return nil
	}
	errPing := dbConn.Ping()
	if errPing != nil {
		log.Error("Error while connecting database", errPing)
		return nil
	}

	log.Info("Successfully connected to sql db...")

	dbConn.SetMaxOpenConns(dbConfigs.MaxPoolSize)
	dbConn.SetMaxIdleConns(dbConfigs.MaxIdleConnections)
	dbConn.SetConnMaxLifetime(time.Minute * 5)
	return dbConn
}
