package routers

import (
	userRepo "GoLangProject/repository"
	userService "GoLangProject/service/User"
	"github.com/gorilla/mux"
	"net/http"
)

func userHandlers(router *mux.Router) {
	router.HandleFunc("/user", createUser).Methods("POST")
	router.HandleFunc("/user/{id:[0-9]+}", updateUser).Methods("PUT")
	router.HandleFunc("/user/{id:[0-9]+}", deleteUser).Methods("DELETE")
	router.HandleFunc("/user/{id:[0-9]+}", selectUser).Methods("GET")
	router.HandleFunc("/user", selectUsers).Methods("GET")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	userRepo := userRepo.UserRepo{}
	userSrv := userService.NewUserService(userRepo)
	sendData(w, userSrv.FindAll(), http.StatusCreated)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	sendData(w, "Update OK", http.StatusOK)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	sendData(w, "Deleted OK", http.StatusOK)
}

func selectUser(w http.ResponseWriter, r *http.Request) {
	sendData(w, "select OK", http.StatusOK)
}

func selectUsers(w http.ResponseWriter, r *http.Request) {
	sendData(w, "select All OK", http.StatusOK)
}
