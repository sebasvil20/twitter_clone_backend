package service

import (
	"github.com/sebasvil20/twitter_clone_backend/utils"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	utils.ResponseMessage(w, "Correct", "Pong", http.StatusOK)
}
