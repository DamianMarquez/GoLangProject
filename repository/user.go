package repository

import "GoLangProject/models"

type UserRepo struct {
}

func (u UserRepo) CreateUser(user *models.User) (*models.User, error) {

	return nil, nil
}

func (u UserRepo) FindAllUsers() []models.User {
	users := []models.User{}
	Database.Find(users, models.User{})
	return users
}

func (u UserRepo) MigrarUser() error {
	return Database.AutoMigrate(models.User{})
}
