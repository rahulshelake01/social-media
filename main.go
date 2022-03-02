package main

import (
	"social-media/config"
	"social-media/server"

	"github.com/joho/godotenv"

	log "github.com/sirupsen/logrus"
)

func init() {

	log.SetFormatter(&log.JSONFormatter{})
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load .env file")
	}
	config.LoadConfig()

}

func main() {
	server := server.Server()
	server.Start()
}
