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

func (m *MockRepository) FindAllUsers() ([]models.User, error) {
	args := m.Called()
	result := args.Get(0)
	return result.([]models.User), args.Error(1)
}

func (m *MockRepository) FindAllTipos() ([]models.Tipo, error) {
	return nil, nil
}

func (m *MockRepository) CreateTipo(user *models.Tipo) (*models.Tipo, error) {
	return nil, nil
}

func (m *MockRepository) FindAllLinks() ([]models.Link, error) {
	return nil, nil
}

func (m *MockRepository) CreateLink(user *models.Link) (*models.Link, error) {
	return nil, nil
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)
	user := models.User{Nombre: "Nombre", Email: "Email"}
	mockRepo.On("FindAllUsers").Return([]models.User{user}, nil)
	userService := NewUserService(mockRepo)
	result, _ := userService.FindAll()
	mockRepo.AssertExpectations(t)
	assert.Equal(t, "Nombre", result[0].Nombre)
	assert.Equal(t, "Email", result[0].Email)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockRepository)
	user := models.User{Nombre: "Nombre", Email: "Email"}
	mockRepo.On("CreateUser").Return(&user, nil)
	userService := NewUserService(mockRepo)
	result, err := userService.Create(&user)
	mockRepo.AssertExpectations(t)
	assert.Equal(t, "Nombre", result.Nombre)
	assert.Equal(t, "Email", result.Email)
	assert.Nil(t, err)
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
