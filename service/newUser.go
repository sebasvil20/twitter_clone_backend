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
		utils.ResponseMessage(w, "HttpMessageNotReadable", "Invalid Email, try again", http.StatusBadRequest)
		return
	}

	if !utils.ValidateEmail(t.Email) {
		utils.ResponseMessage(w, "InvalidEmail", "Invalid Email, try again", http.StatusBadRequest)
		return
	}

	if len(t.Password) < 6 {
		utils.ResponseMessage(w, "InvalidPasswordLenght", "Password minimun lenght is 6 chars", http.StatusBadRequest)
		return
	}

	_, found, _ := repository.FindUserByEmail(t.Email)
	if found {
		utils.ResponseMessage(w, "EmailInUse", "Email is already in use", http.StatusBadRequest)
		return
	}

	saved, err := repository.SaveNewUser(t)
	if err != nil || !saved {
		utils.ResponseMessage(w, "InternalError", "Couldn't save new user, try again later", http.StatusInternalServerError)
		return
	}

	utils.ResponseMessage(w, "Created", "User registered succesfully", http.StatusCreated)
}
