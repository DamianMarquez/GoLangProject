package repository

import "GoLangProject/models"

type MySqlRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	FindAllUsers() ([]models.User, error)

	CreateTipo(user *models.Tipo) (*models.Tipo, error)
	FindAllTipos() ([]models.Tipo, error)

	CreateLink(user *models.Link) (*models.Link, error)
	FindAllLinks() ([]models.Link, error)
}
