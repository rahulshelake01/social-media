package server

import (
	"net/http"
	"os"
	"social-media/api/router"
	"social-media/api/v1/repository"
	"social-media/api/v1/services"
	"social-media/config"
	"social-media/database"

	log "github.com/sirupsen/logrus"
)

type ServerInterface interface {
	Start()
}

type serverStrcut struct{}

func Server() ServerInterface {
	return serverStrcut{}
}

func (s serverStrcut) Start() {

	appConfig := config.GetAppConfig()
	dbConfig := config.GetDBConfig()

	Connection := database.DBConnection(dbConfig)
	db := Connection.DBConnect()

	if db == nil {
		log.Fatal("Expecting db connection object but received nil")
	}

	userRepository := repository.UserRepository(db, appConfig)
	postRepository := repository.SocialMediaRepository(db, appConfig)

	userService := services.UserService(userRepository)
	postService := services.PostService(postRepository)

	router := router.RouteHandler(userService, postService)

	log.Println("Social media server srated on http://" + os.Getenv("SOCIAL_MEDIA_HOST") + ":" + os.Getenv("SOCIAL_MEDIA_PORT"))

	if err := http.ListenAndServe(os.Getenv("SOCIAL_MEDIA_HOST")+":"+os.Getenv("SOCIAL_MEDIA_PORT"), router); err != nil {
		log.Fatal("Failed to start social-media server...")
	}
}
