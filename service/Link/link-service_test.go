package Link

import (
	"GoLangProject/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateLink(user *models.Link) (*models.Link, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*models.Link), args.Error(1)
}

func (m *MockRepository) FindAllLinks() []models.Link {
	args := m.Called()
	result := args.Get(0)
	return result.([]models.Link)
}

func (m *MockRepository) DeleteLink(link *models.Link) (*models.Link, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*models.Link), args.Error(1)
}

func (m *MockRepository) FindLink(id int) models.Link {
	args := m.Called()
	result := args.Get(0)
	return result.(models.Link)
}

func (m *MockRepository) UpdateLink(link *models.Link) (*models.Link, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*models.Link), args.Error(1)
}

func TestFindAllLinks(t *testing.T) {
	mockRepo := new(MockRepository)
	link := models.Link{URL: "URL", Id_user: 1}
	mockRepo.On("FindAllLinks").Return([]models.Link{link}, nil)
	linkService := NewLinkService(mockRepo)
	result := linkService.FindAllLinks()
	mockRepo.AssertExpectations(t)
	assert.Equal(t, "URL", result[0].URL)
}

func TestCreateLink(t *testing.T) {
	mockRepo := new(MockRepository)
	link := models.Link{URL: "URL", Id_user: 1}
	mockRepo.On("CreateLink").Return(&link, nil)
	linkService := NewLinkService(mockRepo)
	result, err := linkService.CreateLink(&link)
	mockRepo.AssertExpectations(t)
	assert.Equal(t, "URL", result.URL)
	assert.Nil(t, err)
}

func TestDeleteLink(t *testing.T) {
	mockRepo := new(MockRepository)
	link := models.Link{URL: "URL", Id_user: 1}
	mockRepo.On("DeleteLink").Return(&link, nil)
	linkService := NewLinkService(mockRepo)
	result, err := linkService.DeleteLink(&link)
	mockRepo.AssertExpectations(t)
	assert.Equal(t, "URL", result.URL)
	assert.Nil(t, err)
}

func TestUdateLink(t *testing.T) {
	mockRepo := new(MockRepository)
	link := models.Link{URL: "URL", Id_user: 1}
	mockRepo.On("UpdateLink").Return(&link, nil)
	linkService := NewLinkService(mockRepo)
	result, err := linkService.UpdateLink(&link)
	mockRepo.AssertExpectations(t)
	assert.Equal(t, "URL", result.URL)
	assert.Nil(t, err)
}

func TestFindLink(t *testing.T) {
	mockRepo := new(MockRepository)
	link := models.Link{URL: "URL", Id_user: 1}
	mockRepo.On("FindLink").Return(link)
	linkService := NewLinkService(mockRepo)
	result := linkService.FindLink(1)
	mockRepo.AssertExpectations(t)
	assert.Equal(t, "URL", result.URL)
	assert.Nil(t, nil)
}

func TestValidateNil(t *testing.T) {
	linkService := NewLinkService(nil)
	err := linkService.Validate(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "The link to validate is nil", err.Error())
}

func TestValidateEmptyURL(t *testing.T) {
	link := models.Link{}
	linkService := NewLinkService(nil)
	err := linkService.Validate(&link)
	assert.NotNil(t, err)
	assert.Equal(t, "The URL of the link it's empty", err.Error())
}

func TestValidateEmptyIdUser(t *testing.T) {
	link := models.Link{URL: "URL"}
	linkService := NewLinkService(nil)
	err := linkService.Validate(&link)
	assert.NotNil(t, err)
	assert.Equal(t, "The User id of the link it's empty", err.Error())
}
