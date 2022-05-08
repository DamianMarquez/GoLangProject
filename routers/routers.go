package routers

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	log.Println("Handlers UP")
	return r
}
