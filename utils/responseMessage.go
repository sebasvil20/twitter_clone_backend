package utils

import (
	"encoding/json"
	"github.com/sebasvil20/twitter_clone_backend/models"
	"net/http"
)

func ResponseMessage(w http.ResponseWriter, code string, inputMessage string, statusCode int) {
	message := models.Message{Code: code, Message: inputMessage}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(message)
}

func ResponseJwtToken(w http.ResponseWriter, jwtToken string, statusCode int) {
	message := models.JwtResponse{JwtToken: jwtToken}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(message)
}
