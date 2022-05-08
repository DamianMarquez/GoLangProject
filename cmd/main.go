package main

import (
	"GoLangProject/routers"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	r := routers.NewRouter()

	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file - ", err)
	}
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "9090"
	}
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Error("Error Crear el Router - ", err)
	}

}
