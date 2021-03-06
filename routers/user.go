package routers

import (
	"GoLangProject/models"
	userRepository "GoLangProject/repository"
	userService "GoLangProject/service/User"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

const usersPath = "/users"

func userHandlers(router *mux.Router) {
	router.HandleFunc(usersPath, createUser).Methods("POST")
	router.HandleFunc(usersPath, updateUser).Methods("PUT")
	router.HandleFunc(usersPath+regExp, deleteUser).Methods("DELETE")
	router.HandleFunc(usersPath+regExp, selectUser).Methods("GET")
	router.HandleFunc(usersPath, selectUsers).Methods("GET")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		log.Println(err)
	} else {
		userRepo := userRepository.UserRepo{}
		userSrv := userService.NewUserService(userRepo)
		if userCreated, err := userSrv.CreateUser(&user); err != nil {
			log.Println("Error creating user: ", err)
		} else {
			sendData(w, userCreated, http.StatusCreated)
		}
	}
	sendError(w, http.StatusBadRequest)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		log.Println(err)
	} else {
		userRepo := userRepository.UserRepo{}
		userSrv := userService.NewUserService(userRepo)
		if userUpdated, err := userSrv.UpdateUser(&user); err != nil {
			log.Println("Error updating user: ", err)
		} else {
			sendData(w, userUpdated, http.StatusOK)
		}
	}
	sendError(w, http.StatusBadRequest)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		log.Println(err)
	} else {
		userRepo := userRepository.UserRepo{}
		userSrv := userService.NewUserService(userRepo)
		if userdeleted, err := userSrv.DeleteUser(&user); err != nil {
			log.Println("Error Deleting user: ", err)
		} else {
			sendData(w, userdeleted, http.StatusOK)
		}
	}
	sendError(w, http.StatusBadRequest)
}

func selectUser(w http.ResponseWriter, r *http.Request) {
	userRepo := userRepository.UserRepo{}
	userSrv := userService.NewUserService(userRepo)
	vars := mux.Vars(r)
	itemId, _ := strconv.Atoi(vars["id"])

	if item := userSrv.FindUser(itemId); item.ID == 0 {
		sendError(w, http.StatusNotFound)
	} else {
		sendData(w, item, http.StatusOK)
	}

}

func selectUsers(w http.ResponseWriter, r *http.Request) {
	userRepo := userRepository.UserRepo{}
	userSrv := userService.NewUserService(userRepo)
	sendData(w, userSrv.FindAllUsers(), http.StatusOK)
}
