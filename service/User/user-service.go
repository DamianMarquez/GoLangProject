package User

import (
	"GoLangProject/models"
	"errors"
)

type UserService interface {
	Validate(user *models.User) error
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(user *models.User) (*models.User, error)
	FindAllUsers() []models.User
	FindUser(id int) models.User
}

var repo MySqlRepository

func NewUserService(repository MySqlRepository) Service {
	repo = repository
	return Service{}
}

type MySqlRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	FindAllUsers() []models.User
	FindUser(int) models.User
	DeleteUser(user *models.User) (*models.User, error)
}

type Service struct {
}

func (*Service) Validate(user *models.User) error {

	if user == nil {
		return errors.New("The user to validate is nil")
	}
	if user.Nombre == "" {
		return errors.New("The Name of the user it's empty")
	}
	if user.Email == "" {
		return errors.New("The Email of the user it's empty")
	}
	return nil
}

func (*Service) CreateUser(user *models.User) (*models.User, error) {
	return repo.CreateUser(user)
}

func (*Service) UpdateUser(user *models.User) (*models.User, error) {
	return repo.UpdateUser(user)
}

func (*Service) FindAllUsers() []models.User {
	return repo.FindAllUsers()
}

func (*Service) FindUser(id int) models.User {
	return repo.FindUser(id)
}

func (*Service) DeleteUser(user *models.User) (*models.User, error) {
	return repo.DeleteUser(user)
}
