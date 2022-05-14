package repository

import "GoLangProject/models"

type TipoRepo struct {
}

func (u TipoRepo) CreateTipo(tipo *models.Tipo) (*models.Tipo, error) {
	if err := Database.Create(&tipo).Error; err != nil {
		return tipo, err
	}
	return tipo, nil
}

func (u TipoRepo) UpdateTipo(tipo *models.Tipo) (*models.Tipo, error) {
	if err := Database.Save(&tipo).Error; err != nil {
		return tipo, err
	}
	return tipo, nil
}

func (u TipoRepo) DeleteTipo(tipo *models.Tipo) (*models.Tipo, error) {
	if err := Database.Delete(&tipo).Error; err != nil {
		return tipo, err
	}
	return tipo, nil
}

func (u TipoRepo) FindAllTipos() []models.Tipo {
	tipos := []models.Tipo{}
	Database.First(&tipos, models.Tipo{})
	return tipos
}

func (u TipoRepo) FindTipo(id int) models.Tipo {
	tipo := models.Tipo{}
	Database.Find(&tipo, id)
	return tipo
}

func (u TipoRepo) MigrarTipo() error {
	return Database.AutoMigrate(models.Tipo{})
}
