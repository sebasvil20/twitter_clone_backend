package service

import (
	"encoding/json"
	"github.com/sebasvil20/twitter_clone_backend/auth"
	"github.com/sebasvil20/twitter_clone_backend/models"
	"github.com/sebasvil20/twitter_clone_backend/repository"
	"github.com/sebasvil20/twitter_clone_backend/utils"
	"net/http"
)

func ModifyUser(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		utils.ResponseMessage(w, "HttpMessageNotReadable", "Couldn't read the http body, try again", http.StatusBadRequest)
		return
	}

	if t.Email != "" {
		utils.ResponseMessage(w, "InvalidChange", "You can't update your email", http.StatusBadRequest)
		return
	}

	var status bool

	status, err = repository.UpdateUser(t, auth.IdUser)
	if err != nil {
		utils.ResponseMessage(w, "InternalError", "Couldn't update user", http.StatusInternalServerError)
		return
	}

	if status {
		w.WriteHeader(http.StatusOK)
	}
}
