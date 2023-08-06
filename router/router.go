package router

import (
	"anonymous-poll/api"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", api.Hello).Methods("GET")
	router.HandleFunc("/getallusers", api.GetAllMyUsers).Methods("GET")
	router.HandleFunc("/register", api.Register).Methods("POST")
	router.HandleFunc("/login", api.Login).Methods("POST")
	router.HandleFunc("/getuser", api.GetOneMyUser).Methods("POST")
	return router

}
