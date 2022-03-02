package services

import (
	"fmt"
	"social-media/models"

	"github.com/stretchr/testify/mock"
)

type mockUserRepository struct {
	m mock.Mock
}

func (m *mockUserRepository) Register(request *models.UserRegisterRequest) error {

	fmt.Printf(".......Mock Register() repository method is called with request : %+v\n", request)
	mockArgs := m.m.Called(request)
	return mockArgs.Error(0)

}

func (m *mockUserRepository) FindUserByEmail(email string) (*models.UserDetails, bool, error) {
	fmt.Printf(".......Mock FindUserByEmail() repository method is called with email : %+v\n", email)
	mockArgs := m.m.Called(email)
	return mockArgs.Get(0).(*models.UserDetails), mockArgs.Get(1).(bool), mockArgs.Error(2)
}
