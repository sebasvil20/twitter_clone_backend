package main

import (
	"github.com/sebasvil20/twitter_clone_backend/database"
	"github.com/sebasvil20/twitter_clone_backend/handlers"
	"log"
)
func main(){
	if database.CheckConnection() == 0 {
		log.Fatal("Can't connect to database")
		return
	}

	handlers.RouterHandlers()
}
