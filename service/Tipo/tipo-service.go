package Tipo

import (
	"GoLangProject/models"
	"GoLangProject/repository"
	"errors"
	"fmt"
)

type TipoService interface {
	Validate(user *models.Tipo) error
	Create(user *models.Tipo) (models.Tipo, error)
	FindAll() ([]models.Tipo, error)
}

type MySqlRepository interface {
	CreateTipo(user *models.Tipo) (*models.Tipo, error)
	FindAllTipos() ([]models.Tipo, error)
}

var repo MySqlRepository

func NewTipoService(repository MySqlRepository) Service {
	repo = repository
	return Service{}
}

type Service struct{}

func (*Service) Validate(tipo *models.Tipo) error {

	if tipo == nil {
		return errors.New("The tipo to validate is nil")
	}
	if tipo.Descripcion == "" {
		return errors.New("The Descripcion of the tipo it's empty")
	}

	return nil
}

func (*Service) Create(tipo *models.Tipo) (*models.Tipo, error) {
	return repo.CreateTipo(tipo)
}

func (*Service) FindAll() ([]models.Tipo, error) {
	return repo.FindAllTipos()
}

func (*Service) MigrarTipo() {
	err := repository.Database.AutoMigrate(models.Tipo{})
	if err != nil {
		fmt.Println("Error al Migrar Tipo: ", err)
	}
}
