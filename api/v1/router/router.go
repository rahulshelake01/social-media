package router

import (
	"encoding/json"
	"net/http"
	"social-media/api/v1/services"
	"social-media/api/v1/utils"
	"social-media/models"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type RouterStruct struct {
	User        services.UserServiceInterface
	PostService services.PostServiceInterfce
}

func RouteHandler(userService services.UserServiceInterface, postService services.PostServiceInterfce) *mux.Router {
	handler := &RouterStruct{User: userService, PostService: postService}

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/user/register", handler.Register).Methods("POST")
	router.HandleFunc("/api/v1/user/login", handler.Login).Methods("POST")
	router.HandleFunc("/api/v1/post/add", handler.AddPost).Methods("POST")

	return router
}

func (ph RouterStruct) Register(w http.ResponseWriter, r *http.Request) {

	var request models.UserRegisterRequest
	json.NewDecoder(r.Body).Decode(&request)

	log.Info("Register request --- ", request)

	response := ph.User.Register(&request)
	utils.ResponseJson(w, response)

}

func (ph RouterStruct) Login(w http.ResponseWriter, r *http.Request) {

	var request models.UserLoginRequest
	json.NewDecoder(r.Body).Decode(&request)

	log.Info("Register request --- ", request)

	response := ph.User.Login(&request)
	utils.ResponseJson(w, response)

}

func (ph RouterStruct) AddPost(w http.ResponseWriter, r *http.Request) {
	utils.ResponseJson(w, "Post added successfully....")
}
