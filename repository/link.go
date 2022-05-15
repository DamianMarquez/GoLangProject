package repository

import "GoLangProject/models"

type LinkRepo struct {
}

func (u LinkRepo) CreateLink(Link *models.Link) (*models.Link, error) {
	if err := Database.Create(&Link).Error; err != nil {
		return Link, err
	}
	return Link, nil
}

func (u LinkRepo) UpdateLink(Link *models.Link) (*models.Link, error) {
	if err := Database.Save(&Link).Error; err != nil {
		return Link, err
	}
	return Link, nil
}

func (u LinkRepo) DeleteLink(Link *models.Link) (*models.Link, error) {
	if err := Database.Delete(&Link).Error; err != nil {
		return Link, err
	}
	return Link, nil
}

func (u LinkRepo) FindAllLinks() []models.Link {
	Links := []models.Link{}
	Database.First(&Links, models.Link{})
	return Links
}

func (u LinkRepo) FindLink(id int) models.Link {
	Link := models.Link{}
	Database.Find(&Link, id)
	return Link
}

func (u LinkRepo) MigrarLink() error {
	return Database.AutoMigrate(models.Link{})
}
