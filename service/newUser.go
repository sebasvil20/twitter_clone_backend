package service

import (
	"encoding/json"
	"github.com/sebasvil20/twitter_clone_backend/models"
	"github.com/sebasvil20/twitter_clone_backend/repository"
	"github.com/sebasvil20/twitter_clone_backend/utils"
	"net/http"
)

func RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var t models.User
	
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "HttpMessageNotReadable", http.StatusBadRequest)
		return
	}

	if !utils.ValidateEmail(t.Email) {
		http.Error(w, "Invalid Email, try again", http.StatusBadRequest)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password minimun lenght is 6 chars", http.StatusBadRequest)
		return
	}

	_, found, _ := repository.FindUserByEmail(t.Email)
	if found {
		http.Error(w, "Email is already in use", http.StatusBadRequest)
		return
	}

	saved, err := repository.SaveNewUser(t)
	if err != nil || !saved {
		http.Error(w, "Couldn't save new user, try again later", http.StatusInternalServerError)
		return
	}

	message := models.Message{Message: "User registered succesfully"}
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		http.Error(w, "Couldn't format to json", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err2 := w.Write(jsonMessage)
	if err2 != nil {
		return
	}

}
