package repository

import (
	"GoLangProject/models"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

//Realiza la conexion

var Database = func() (db *gorm.DB) {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")

	var dsn = user + ":" + pass + host + name + "?charset=utf8mb4&parseTime=True&loc=Local"

	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		log.Println("Error en la conexion", err)
		panic(err)
	} else {
		log.Error("Conexion exitosa")
		return db
	}
}()

func Migrar() error {
	return Database.AutoMigrate(models.User{})
}
