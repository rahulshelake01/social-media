package services

import (
	"fmt"
	"social-media/api/v1/repository"
	"social-media/api/v1/utils"
	"social-media/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("MY_JWT_SECRET")

type Claims struct {
	Uid   int64  `json:"uid"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type UserServiceInterface interface {
	Register(request *models.UserRegisterRequest) *models.UserRegisterResponse
	Login(request *models.UserLoginRequest) *models.UserLoginResponse
}

type UserServiceStruct struct {
	UserRepository repository.UserRepositoryInterface
}

func UserService(userRepo repository.UserRepositoryInterface) UserServiceInterface {
	return UserServiceStruct{UserRepository: userRepo}
}

func (userService UserServiceStruct) Register(request *models.UserRegisterRequest) *models.UserRegisterResponse {

	var (
		response = new(models.UserRegisterResponse)
		err      error
	)

	user, isFound, err := userService.UserRepository.FindUserByEmail(request.Email)
	if err != nil {
		fmt.Println("Failed to fetch user by email")
		response.Message = "Something went wrong."
		response.Success = false
		response.StatusCode = 300
		return response
	}

	if isFound {
		fmt.Println("User alredy registered.")
		fmt.Println(user)
		response.Message = "User alredy registered."
		response.Success = false
		response.StatusCode = 300
		return response
	}

	request.Password, err = utils.HashPassword(request.Password)

	if err != nil {
		fmt.Println("Failed to hash password.")
		response.Message = "Something went wrong."
		return response
	}

	err = userService.UserRepository.Register(request)

	if err != nil {
		fmt.Println("Failed to register user.")
		response.Message = "Something went wrong."
		response.Success = false
		response.StatusCode = 300
		return response
	}

	response.Message = "User registered successfully."
	response.Success = true
	response.StatusCode = 200
	return response
}

func (userService UserServiceStruct) Login(request *models.UserLoginRequest) *models.UserLoginResponse {

	var (
		response = new(models.UserLoginResponse)
		err      error
	)

	user, isFound, err := userService.UserRepository.FindUserByEmail(request.Email)

	fmt.Printf("****************************")
	fmt.Printf(" user : %+v\n", user)

	if err != nil {
		fmt.Println("Failed to fetch user by email")
		response.Message = "Something went wrong."
		return response
	}

	if !isFound {
		fmt.Println("User not found.")
		fmt.Println(user)
		response.Message = "User not found."
		return response
	}

	if !utils.CheckPasswordHash(request.Password, user.Password) {
		fmt.Println("Wrong passwird.")
		fmt.Println(user)
		response.Message = "Invalid creds."
		return response
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := Claims{
		Uid:   user.UID,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println("Failed to create JWT token")
		response.Message = "Something went wrong."
		return response
	}

	response.Success = true
	response.Data.Email = user.Email
	response.Data.Token = tokenString
	return response
}
