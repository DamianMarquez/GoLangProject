package Tipo

import (
	"GoLangProject/models"
	"GoLangProject/repository"
	"errors"
	"fmt"
)

type TipoService interface {
	Validate(user *models.Tipo) error
	CreateTipo(user *models.Tipo) (*models.Tipo, error)
	UpdateTipo(user *models.Tipo) (*models.Tipo, error)
	DeleteTipo(user *models.Tipo) (*models.Tipo, error)
	FindAllTipos() []models.Tipo
	FindTipo(id int) models.Tipo
}

type MySqlRepository interface {
	CreateTipo(user *models.Tipo) (*models.Tipo, error)
	UpdateTipo(user *models.Tipo) (*models.Tipo, error)
	DeleteTipo(user *models.Tipo) (*models.Tipo, error)
	FindAllTipos() []models.Tipo
	FindTipo(id int) models.Tipo
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

func (*Service) CreateTipo(tipo *models.Tipo) (*models.Tipo, error) {
	return repo.CreateTipo(tipo)
}

func (*Service) UpdateTipo(tipo *models.Tipo) (*models.Tipo, error) {
	return repo.UpdateTipo(tipo)
}

func (*Service) FindAllTipos() []models.Tipo {
	return repo.FindAllTipos()
}

func (*Service) FindTipo(id int) models.Tipo {
	return repo.FindTipo(id)
}

func (*Service) DeleteTipo(tipo *models.Tipo) (*models.Tipo, error) {
	return repo.DeleteTipo(tipo)
}

func (*Service) MigrarTipo() {
	err := repository.Database.AutoMigrate(models.Tipo{})
	if err != nil {
		fmt.Println("Error al Migrar Tipo: ", err)
	}
}
