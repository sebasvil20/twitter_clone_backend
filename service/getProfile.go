package service

import (
	"encoding/json"
	"github.com/sebasvil20/twitter_clone_backend/repository"
	"github.com/sebasvil20/twitter_clone_backend/utils"
	"net/http"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		utils.ResponseMessage(w, "MandatoryParameter", "Id Parameter is mandaroty", http.StatusBadRequest)
		return
	}

	profile, err := repository.FindProfileByID(ID)

	if err != nil {
		utils.ResponseMessage(w, "UserNotFound", "User with wanted id doesn't exists", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}
