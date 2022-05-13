package Tipo

import (
	"GoLangProject/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateTipo(user *models.Tipo) (*models.Tipo, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*models.Tipo), args.Error(1)
}

func (m *MockRepository) FindAllTipos() []models.Tipo {
	args := m.Called()
	result := args.Get(0)
	return result.([]models.Tipo)
}

func (m *MockRepository) DeleteTipo(user *models.Tipo) (*models.Tipo, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*models.Tipo), args.Error(1)
}

func (m *MockRepository) FindTipo(id int) models.Tipo {
	args := m.Called()
	result := args.Get(0)
	return result.(models.Tipo)
}

func (m *MockRepository) UpdateTipo(user *models.Tipo) (*models.Tipo, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*models.Tipo), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)
	tipo := models.Tipo{Descripcion: "Descripcion"}
	mockRepo.On("FindAllTipos").Return([]models.Tipo{tipo}, nil)
	tipoService := NewTipoService(mockRepo)
	result := tipoService.FindAllTipos()
	mockRepo.AssertExpectations(t)
	assert.Equal(t, "Descripcion", result[0].Descripcion)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockRepository)
	tipo := models.Tipo{Descripcion: "Descripcion"}
	mockRepo.On("CreateTipo").Return(&tipo, nil)
	tipoService := NewTipoService(mockRepo)
	result, err := tipoService.CreateTipo(&tipo)
	mockRepo.AssertExpectations(t)
	assert.Equal(t, "Descripcion", result.Descripcion)
	assert.Nil(t, err)
}

func TestValidateNil(t *testing.T) {
	tipoService := NewTipoService(nil)
	err := tipoService.Validate(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "The tipo to validate is nil", err.Error())
}

func TestValidateEmptyDescription(t *testing.T) {
	tipo := models.Tipo{}
	tipoService := NewTipoService(nil)
	err := tipoService.Validate(&tipo)
	assert.NotNil(t, err)
	assert.Equal(t, "The Descripcion of the tipo it's empty", err.Error())
}
