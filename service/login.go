package service

import (
	"encoding/json"
	"github.com/sebasvil20/twitter_clone_backend/auth"
	"github.com/sebasvil20/twitter_clone_backend/models"
	"github.com/sebasvil20/twitter_clone_backend/repository"
	"github.com/sebasvil20/twitter_clone_backend/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		utils.ResponseMessage(w, "HttpMessageNotReadable", "Couldn't read the http body, try again", http.StatusBadRequest)
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

	document, exists := compareCredentials(t.Email, t.Password)

	if !exists {
		utils.ResponseMessage(w, "BadCredentials", "Invalid email or password", http.StatusBadRequest)
		return
	}

	jwtKey, err := auth.GenerateJWT(document)
	if err != nil {
		utils.ResponseMessage(w, "InternalError", "Couldn't generate jwt token, try again later", http.StatusInternalServerError)
		return
	}

	utils.ResponseJwtToken(w, jwtKey, http.StatusOK)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}

func compareCredentials(email string, password string) (models.User, bool) {
	user, found, _ := repository.FindUserByEmail(email)
	if !found {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDb := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordDb, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
