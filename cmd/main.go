package main

import (
	"GoLangProject/repository"
	"GoLangProject/routers"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {

	var port string

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file - ", err)
	} else {
		port = os.Getenv("API_PORT")
	}

	if port == "" {
		port = "9090" //DEFAULT PORT
	}

	if err := repository.Migrar(); err != nil {
		log.Fatal("Error Migrating models - ", err)
	}

	r := routers.NewRouter()

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Error("Error Crear el Router - ", err)
	}

}
