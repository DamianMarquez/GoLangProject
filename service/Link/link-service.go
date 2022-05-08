package Link

import (
	"GoLangProject/models"
	"GoLangProject/repository"
	"errors"
	"fmt"
)

type LinkService interface {
	Validate(user *models.Link) error
	Create(user *models.Link) (models.Link, error)
	FindAll() ([]models.Link, error)
}

var repo repository.MySqlRepository

func NewLinkService(repository repository.MySqlRepository) Service {
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

func (*Service) Create(link *models.Link) (*models.Link, error) {
	return repo.CreateLink(link)
}

func (*Service) FindAll() ([]models.Link, error) {
	return repo.FindAllLinks()
}

func (*Service) MigrarLink() {
	err := repository.Database.AutoMigrate(models.Link{})
	if err != nil {
		fmt.Println("Error al Migrar Link: ", err)
	}
}
