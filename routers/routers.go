package routers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const regExp = "/{id:[0-9]+}"

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	userHandlers(r)
	tipoHandlers(r)
	linkHandlers(r)

	log.Println("Handlers UP")
	return r
}

func sendData(rw http.ResponseWriter, data interface{}, status int) {
	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(status)
	err := json.NewEncoder(rw).Encode(data)
	if err != nil {
		return
	}
	output, _ := json.Marshal(&data)
	log.Println(rw, string(output))

}

func sendError(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)
	log.Println(rw, "Resource Not Found")
}
