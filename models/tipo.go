package models

import "gorm.io/gorm"

type Tipo struct {
	gorm.Model
	Descripcion string
}
