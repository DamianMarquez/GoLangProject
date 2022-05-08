package routers

import (
	"GoLangProject/models"
	userRepository "GoLangProject/repository"
	userService "GoLangProject/service/User"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func userHandlers(router *mux.Router) {
	router.HandleFunc("/user", createUser).Methods("POST")
	router.HandleFunc("/user", updateUser).Methods("PUT")
	router.HandleFunc("/user/{id:[0-9]+}", deleteUser).Methods("DELETE")
	router.HandleFunc("/user/{id:[0-9]+}", selectUser).Methods("GET")
	router.HandleFunc("/user", selectUsers).Methods("GET")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		log.Println(err)
	} else {
		userRepo := userRepository.UserRepo{}
		userSrv := userService.NewUserService(userRepo)
		if userCreated, err := userSrv.Create(&user); err != nil {
			log.Println("Error creating user: ", err)
		} else {
			sendData(w, userCreated, http.StatusCreated)
		}
	}
	sendError(w, http.StatusInternalServerError)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		log.Println(err)
	} else {
		userRepo := userRepository.UserRepo{}
		userSrv := userService.NewUserService(userRepo)
		if userUpdated, err := userSrv.Update(&user); err != nil {
			log.Println("Error updating user: ", err)
		} else {
			sendData(w, userUpdated, http.StatusCreated)
		}
	}
	sendError(w, http.StatusInternalServerError)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	sendData(w, "Deleted OK", http.StatusOK)
}

func selectUser(w http.ResponseWriter, r *http.Request) {
	sendData(w, "select OK", http.StatusOK)
}

func selectUsers(w http.ResponseWriter, r *http.Request) {
	userRepo := userRepository.UserRepo{}
	userSrv := userService.NewUserService(userRepo)
	sendData(w, userSrv.FindAll(), http.StatusCreated)

}
