package User

import (
	"GoLangProject/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateUser(user *models.User) (*models.User, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*models.User), args.Error(1)
}

func (m *MockRepository) FindAllUsers() []models.User {
	args := m.Called()
	result := args.Get(0)
	return result.([]models.User)
}

func (m *MockRepository) DeleteUser(user *models.User) (*models.User, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*models.User), args.Error(1)
}

func (m *MockRepository) FindUser(id int) models.User {
	args := m.Called()
	result := args.Get(0)
	return result.(models.User)
}

func (m *MockRepository) UpdateUser(user *models.User) (*models.User, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*models.User), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockRepository)
	user := models.User{Nombre: "Nombre", Email: "Email"}
	mockRepo.On("CreateUser").Return(&user, nil)
	userService := NewUserService(mockRepo)
	result, err := userService.CreateUser(&user)
	mockRepo.AssertExpectations(t)
	assert.Equal(t, "Nombre", result.Nombre)
	assert.Equal(t, "Email", result.Email)
	assert.Nil(t, err)
}

func TestFindAllUsers(t *testing.T) {
	mockRepo := new(MockRepository)
	user := models.User{Nombre: "Nombre", Email: "Email"}
	mockRepo.On("FindAllUsers").Return([]models.User{user}, nil)
	userService := NewUserService(mockRepo)
	result := userService.FindAllUsers()
	mockRepo.AssertExpectations(t)
	assert.Equal(t, "Nombre", result[0].Nombre)
	assert.Equal(t, "Email", result[0].Email)
}

func TestDeleteUser(t *testing.T) {
	mockRepo := new(MockRepository)
	user := models.User{Nombre: "Nombre", Email: "Email"}
	mockRepo.On("DeleteUser").Return(&user, nil)
	userService := NewUserService(mockRepo)
	result, err := userService.DeleteUser(&user)
	mockRepo.AssertExpectations(t)
	assert.Equal(t, "Nombre", result.Nombre)
	assert.Equal(t, "Email", result.Email)
	assert.Nil(t, err)
}

func TestFindUser(t *testing.T) {
	mockRepo := new(MockRepository)
	user := models.User{Nombre: "Nombre", Email: "Email"}
	mockRepo.On("FindUser").Return(user)
	userService := NewUserService(mockRepo)
	result := userService.FindUser(1)
	mockRepo.AssertExpectations(t)
	assert.Equal(t, "Nombre", result.Nombre)
	assert.Equal(t, "Email", result.Email)
	assert.Nil(t, nil)
}

func TestValidateNil(t *testing.T) {
	userService := NewUserService(nil)
	err := userService.Validate(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "The user to validate is nil", err.Error())
}

func TestValidateEmptyNombre(t *testing.T) {
	user := models.User{Nombre: ""}
	userService := NewUserService(nil)
	err := userService.Validate(&user)
	assert.NotNil(t, err)
	assert.Equal(t, "The Name of the user it's empty", err.Error())
}

func TestValidateEmptyEmail(t *testing.T) {
	user := models.User{Nombre: "Nombre", Email: ""}
	userService := NewUserService(nil)
	err := userService.Validate(&user)
	assert.NotNil(t, err)
	assert.Equal(t, "The Email of the user it's empty", err.Error())
}
