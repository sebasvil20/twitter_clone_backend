package controller

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sebasvil20/twitter_clone_backend/middlewares"
	"github.com/sebasvil20/twitter_clone_backend/service"
	"github.com/sebasvil20/twitter_clone_backend/utils"
	"log"
	"net/http"
)

//RouterHandlers handle routes and server
func RouterHandlers() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", service.HelloWorld).Methods("GET")
	router.HandleFunc("/users/register", middlewares.CheckDB(service.RegisterNewUser)).Methods("POST")
	router.HandleFunc("/users/login", middlewares.CheckDB(service.Login)).Methods("POST")
	router.HandleFunc("/users/profile", middlewares.CheckDB(middlewares.ValidateJWT(service.GetProfile))).Methods("GET")
	router.HandleFunc("/users/update", middlewares.CheckDB(middlewares.ValidateJWT(service.ModifyUser))).Methods("PUT")

	PORT := utils.GotEnvVariable("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
