package service

import (
	"encoding/json"
	"github.com/sebasvil20/twitter_clone_backend/repository"
	"github.com/sebasvil20/twitter_clone_backend/utils"
	"net/http"
	"strconv"
)

func TweetsByUserId(w http.ResponseWriter, r * http.Request) {

	id := r.URL.Query().Get("id")
	if len(id) < 1 {
		utils.ResponseMessage(w, "MandatoryParameter", "User id is mandatory", http.StatusBadRequest)
		return
	}
	pageParameter := r.URL.Query().Get("page")
	if len(pageParameter) < 1 {
		utils.ResponseMessage(w, "MandatoryParameter", "Page number is mandatory", http.StatusBadRequest)
	}

	pageNumber, err := strconv.Atoi(pageParameter)
	if err != nil {
		utils.ResponseMessage(w, "FormatError", "Cannot convert page to a number, try again", http.StatusBadRequest)
	}

	page := int64(pageNumber)
	response, found := repository.GetTweetsByUserId(id, page)
	if !found {
		utils.ResponseMessage(w, "InternalError", "Couldn't find tweets, try again later", http.StatusBadRequest)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}