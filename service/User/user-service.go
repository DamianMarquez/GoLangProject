package User

import (
	"GoLangProject/models"
	"GoLangProject/repository"
	"errors"
	"fmt"
)

type UserService interface {
	Validate(user *models.User) error
	Create(user *models.User) (models.User, error)
	FindAll() ([]models.User, error)
}

var repo MySqlRepository

func NewUserService(repository MySqlRepository) Service {
	repo = repository
	return Service{}
}

type MySqlRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	FindAllUsers() ([]models.User, error)
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

func (*Service) Create(user *models.User) (*models.User, error) {
	return repo.CreateUser(user)
}

func (*Service) FindAll() ([]models.User, error) {
	return repo.FindAllUsers()
}

func (*Service) MigrarUser() {
	err := repository.Database.AutoMigrate(models.User{})
	if err != nil {
		fmt.Println("Error al Migrar User: ", err)
	}
}
