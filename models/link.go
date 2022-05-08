package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	URL     string
	Id_user uint
}
