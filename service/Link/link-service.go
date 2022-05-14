package Link

import (
	"GoLangProject/models"
	"GoLangProject/repository"
	"errors"
	"fmt"
)

type LinkService interface {
	Validate(user *models.Link) error
	CreateLink(user *models.Link) (*models.Link, error)
	UpdateLink(user *models.Link) (*models.Link, error)
	DeleteLink(user *models.Link) (*models.Link, error)
	FindAllLinks() []models.Link
	FindLink(id int) models.Link
}

type MySqlRepository interface {
	CreateLink(user *models.Link) (*models.Link, error)
	UpdateLink(user *models.Link) (*models.Link, error)
	DeleteLink(user *models.Link) (*models.Link, error)
	FindAllLinks() []models.Link
	FindLink(id int) models.Link
}

var repo MySqlRepository

func NewLinkService(repository MySqlRepository) Service {
	repo = repository
	return Service{}
}

type Service struct{}

func (*Service) Validate(link *models.Link) error {

	if link == nil {
		return errors.New("The link to validate is nil")
	}
	if link.URL == "" {
		return errors.New("The URL of the link it's empty")
	}
	if link.Id_user == 0 {
		return errors.New("The User id of the link it's empty")
	}
	return nil
}

func (*Service) CreateLink(link *models.Link) (*models.Link, error) {
	return repo.CreateLink(link)
}

func (*Service) FindAllLinks() []models.Link {
	return repo.FindAllLinks()
}

func (*Service) UpdateLink(link *models.Link) (*models.Link, error) {
	return repo.UpdateLink(link)
}

func (*Service) FindLink(id int) models.Link {
	return repo.FindLink(id)
}

func (*Service) DeleteLink(tipo *models.Link) (*models.Link, error) {
	return repo.DeleteLink(tipo)
}

func (*Service) MigrarLink() {
	err := repository.Database.AutoMigrate(models.Link{})
	if err != nil {
		fmt.Println("Error al Migrar Link: ", err)
	}
}
