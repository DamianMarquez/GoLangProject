package routers

import (
	"GoLangProject/models"
	linkRepository "GoLangProject/repository"
	linkService "GoLangProject/service/Link"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func linkHandlers(router *mux.Router) {
	router.HandleFunc("/links", createLink).Methods("POST")
	router.HandleFunc("/links", updateLink).Methods("PUT")
	router.HandleFunc("/links/{id:[0-9]+}", deleteLink).Methods("DELETE")
	router.HandleFunc("/links/{id:[0-9]+}", selectLink).Methods("GET")
	router.HandleFunc("/links", selectLinks).Methods("GET")
}

func createLink(w http.ResponseWriter, r *http.Request) {
	Link := models.Link{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&Link); err != nil {
		log.Println(err)
	} else {
		LinkRepo := linkRepository.LinkRepo{}
		LinkSrv := linkService.NewLinkService(LinkRepo)
		if LinkCreated, err := LinkSrv.CreateLink(&Link); err != nil {
			log.Println("Error creating Link: ", err)
		} else {
			sendData(w, LinkCreated, http.StatusCreated)
		}
	}
	sendError(w, http.StatusBadRequest)
}

func updateLink(w http.ResponseWriter, r *http.Request) {
	Link := models.Link{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&Link); err != nil {
		log.Println(err)
	} else {
		LinkRepo := linkRepository.LinkRepo{}
		LinkSrv := linkService.NewLinkService(LinkRepo)
		if LinkUpdated, err := LinkSrv.UpdateLink(&Link); err != nil {
			log.Println("Error updating Link: ", err)
		} else {
			sendData(w, LinkUpdated, http.StatusOK)
		}
	}
	sendError(w, http.StatusBadRequest)
}

func deleteLink(w http.ResponseWriter, r *http.Request) {
	Link := models.Link{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&Link); err != nil {
		log.Println(err)
	} else {
		LinkRepo := linkRepository.LinkRepo{}
		LinkSrv := linkService.NewLinkService(LinkRepo)
		if Linkdeleted, err := LinkSrv.DeleteLink(&Link); err != nil {
			log.Println("Error Deleting Link: ", err)
		} else {
			sendData(w, Linkdeleted, http.StatusOK)
		}
	}
	sendError(w, http.StatusBadRequest)
}

func selectLink(w http.ResponseWriter, r *http.Request) {
	LinkRepo := linkRepository.LinkRepo{}
	LinkSrv := linkService.NewLinkService(LinkRepo)
	vars := mux.Vars(r)
	itemId, _ := strconv.Atoi(vars["id"])

	if item := LinkSrv.FindLink(itemId); item.ID == 0 {
		sendError(w, http.StatusNotFound)
	} else {
		sendData(w, item, http.StatusOK)
	}

}

func selectLinks(w http.ResponseWriter, r *http.Request) {
	LinkRepo := linkRepository.LinkRepo{}
	LinkSrv := linkService.NewLinkService(&LinkRepo)
	sendData(w, LinkSrv.FindAllLinks(), http.StatusOK)
}
