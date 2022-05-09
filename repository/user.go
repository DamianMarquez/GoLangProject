package repository

import "GoLangProject/models"

type UserRepo struct {
}

func (u UserRepo) CreateUser(user *models.User) (*models.User, error) {
	if err := Database.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u UserRepo) UpdateUser(user *models.User) (*models.User, error) {
	if err := Database.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u UserRepo) DeleteUser(user *models.User) (*models.User, error) {
	if err := Database.Delete(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u UserRepo) FindAllUsers() []models.User {
	users := []models.User{}
	Database.First(&users, models.User{})
	return users
}

func (u UserRepo) FindUser(id int) models.User {
	user := models.User{}
	Database.Find(&user, id)
	return user
}

func (u UserRepo) MigrarUser() error {
	return Database.AutoMigrate(models.User{})
}
