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
	apiPrefix := "api/v1"

	router.HandleFunc(apiPrefix + "/register", middlewares.CheckDB(service.RegisterNewUser)).Methods("POST")

	PORT := utils.GotEnvVariable("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
