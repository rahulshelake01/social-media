package services

import (
	"errors"
	"social-media/models"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type userSuite struct {
	suite.Suite
	mockUserRepository *mockUserRepository
	userService        UserServiceInterface
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(userSuite))
}

func (u *userSuite) SetupTest() {

	u.mockUserRepository = new(mockUserRepository)
	u.userService = UserService(u.mockUserRepository)

}

func (u *userSuite) TestRegisterWithUserAlreadyExistCase() {

	var (
		request          = getUserRegisterRequest()
		expectedResponse = getUserAlreadyRegisterdResponse()
	)

	u.mockUserRepository.m.On("FindUserByEmail", mock.Anything).Return(getFindUserByEmailResponse(), true, nil)

	response := u.userService.Register(request)
	u.Equal(expectedResponse, response)

	// u.mockUserRepository.m.AssertExpectations(u.Suite.T())

}

func (u *userSuite) TestRegisterWithFail() {

	var (
		request          = getUserRegisterRequest()
		expectedResponse = getUserRegisterFailResponse()
	)

	u.mockUserRepository.m.On("FindUserByEmail", mock.Anything).Return(getFindUserByEmailResponse(), false, nil)
	u.mockUserRepository.m.On("Register", mock.Anything).Return(errors.New("Failed to login"))

	response := u.userService.Register(request)
	u.Equal(expectedResponse, response)

	// u.mockUserRepository.m.AssertExpectations(u.Suite.T())

}

func (u *userSuite) TestRegisterWithSuccess() {

	var (
		request          = getUserRegisterRequest()
		expectedResponse = getUserRegisterResponse()
	)

	u.mockUserRepository.m.On("FindUserByEmail", mock.Anything).Return(getFindUserByEmailResponse(), false, nil)
	u.mockUserRepository.m.On("Register", mock.Anything).Return(nil)

	response := u.userService.Register(request)
	u.Equal(expectedResponse, response)

	// u.mockUserRepository.m.AssertExpectations(u.Suite.T())

}

func getUserRegisterRequest() *models.UserRegisterRequest {

	return &models.UserRegisterRequest{
		FirstName: "RahulTest",
		LastName:  "ShelakeTest",
		Email:     "rahult@test.com",
		Password:  "abctest",
	}
}

func getUserAlreadyRegisterdResponse() *models.UserRegisterResponse {
	return &models.UserRegisterResponse{
		Success:    false,
		StatusCode: 300,
		Message:    "User alredy registered.",
	}
}

func getUserRegisterFailResponse() *models.UserRegisterResponse {
	return &models.UserRegisterResponse{
		Success:    false,
		StatusCode: 300,
		Message:    "Something went wrong.",
	}
}

func getUserRegisterResponse() *models.UserRegisterResponse {
	return &models.UserRegisterResponse{
		Success:    true,
		StatusCode: 200,
		Message:    "User registered successfully.",
	}
}

func getFindUserByEmailResponse() *models.UserDetails {

	return &models.UserDetails{
		UID:       1,
		FirstName: "RahulTest",
		LastName:  "ShelakeTest",
		Email:     "rahult@test.com",
		Password:  "abctest",
	}

}
