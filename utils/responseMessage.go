package utils

import (
	"encoding/json"
	"github.com/sebasvil20/twitter_clone_backend/models"
	"net/http"
)

func ResponseMessage(w http.ResponseWriter, code string, inputMessage string, statusCode int) {
	message := models.Message{Code: code, Message: inputMessage}
	jsonMessage, _ := json.Marshal(message)
	w.WriteHeader(statusCode)
	_, err2 := w.Write(jsonMessage)
	if err2 != nil {
		return
	}
}
