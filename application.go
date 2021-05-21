package main

import (
	"log"
	"net/http"

	"github.com/aadesh-agarwal8888/API-Service/configs"
	"github.com/aadesh-agarwal8888/API-Service/handler"
	"github.com/gorilla/mux"
)

func main() {
	configuration, err := configs.LoadConfigurations()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/login", handler.Login).Methods(http.MethodPost)           //Login => Post
	router.HandleFunc("/register", handler.RegisterUser).Methods(http.MethodPost) //Register => Post
	router.HandleFunc("/user", handler.GetUserDetails).Methods(http.MethodGet)    //Get User Data

	server := http.Server{
		Addr:    configuration.Api_Service,
		Handler: router,
	}

	log.Println("Starting API-Service on " + configuration.Api_Service)
	log.Fatal(server.ListenAndServe())

}
