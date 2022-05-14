package routers

import (
	"GoLangProject/models"
	tipoRepository "GoLangProject/repository"
	tipoService "GoLangProject/service/Tipo"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func tipoHandlers(router *mux.Router) {
	router.HandleFunc("/tipos", createTipo).Methods("POST")
	router.HandleFunc("/tipos", updateTipo).Methods("PUT")
	router.HandleFunc("/tipos/{id:[0-9]+}", deleteTipo).Methods("DELETE")
	router.HandleFunc("/tipos/{id:[0-9]+}", selectTipo).Methods("GET")
	router.HandleFunc("/tipos", selectTipos).Methods("GET")
}

func createTipo(w http.ResponseWriter, r *http.Request) {
	tipo := models.Tipo{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tipo); err != nil {
		log.Println(err)
	} else {
		tipoRepo := tipoRepository.TipoRepo{}
		tipoSrv := tipoService.NewTipoService(tipoRepo)
		if tipoCreated, err := tipoSrv.CreateTipo(&tipo); err != nil {
			log.Println("Error creating tipo: ", err)
		} else {
			sendData(w, tipoCreated, http.StatusCreated)
		}
	}
	sendError(w, http.StatusBadRequest)
}

func updateTipo(w http.ResponseWriter, r *http.Request) {
	tipo := models.Tipo{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tipo); err != nil {
		log.Println(err)
	} else {
		tipoRepo := tipoRepository.TipoRepo{}
		tipoSrv := tipoService.NewTipoService(tipoRepo)
		if tipoUpdated, err := tipoSrv.UpdateTipo(&tipo); err != nil {
			log.Println("Error updating tipo: ", err)
		} else {
			sendData(w, tipoUpdated, http.StatusOK)
		}
	}
	sendError(w, http.StatusBadRequest)
}

func deleteTipo(w http.ResponseWriter, r *http.Request) {
	tipo := models.Tipo{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tipo); err != nil {
		log.Println(err)
	} else {
		tipoRepo := tipoRepository.TipoRepo{}
		tipoSrv := tipoService.NewTipoService(tipoRepo)
		if tipodeleted, err := tipoSrv.DeleteTipo(&tipo); err != nil {
			log.Println("Error Deleting tipo: ", err)
		} else {
			sendData(w, tipodeleted, http.StatusOK)
		}
	}
	sendError(w, http.StatusBadRequest)
}

func selectTipo(w http.ResponseWriter, r *http.Request) {
	tipoRepo := tipoRepository.TipoRepo{}
	tipoSrv := tipoService.NewTipoService(tipoRepo)
	vars := mux.Vars(r)
	itemId, _ := strconv.Atoi(vars["id"])

	if item := tipoSrv.FindTipo(itemId); item.ID == 0 {
		sendError(w, http.StatusNotFound)
	} else {
		sendData(w, item, http.StatusOK)
	}

}

func selectTipos(w http.ResponseWriter, r *http.Request) {
	tipoRepo := tipoRepository.TipoRepo{}
	tipoSrv := tipoService.NewTipoService(&tipoRepo)
	sendData(w, tipoSrv.FindAllTipos(), http.StatusOK)
}
