package service

import (
	"encoding/json"
	"github.com/sebasvil20/twitter_clone_backend/auth"
	"github.com/sebasvil20/twitter_clone_backend/models"
	"github.com/sebasvil20/twitter_clone_backend/repository"
	"github.com/sebasvil20/twitter_clone_backend/utils"
	"net/http"
	"time"
)

func AddNewTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tweet models.Tweet

	err := json.NewDecoder(r.Body).Decode(&tweet)
	if err != nil {
		utils.ResponseMessage(w, "HttpMessageNotReadable", "Couldn'tweet read the http body, try again", http.StatusBadRequest)
		return
	}

	document := models.Tweet{
		UserId:  auth.IdUser,
		Content: tweet.Content,
		Date:    time.Now(),
	}

	if len(tweet.Content) < 1 {
		utils.ResponseMessage(w, "InvalidTweetLength", "Tweet content minimun lenght is 1 chars", http.StatusBadRequest)
		return
	}

	_, saved, err := repository.SaveNewTweet(document)
	if err != nil || !saved {
		utils.ResponseMessage(w, "InternalError", "Couldn't save tweet, try again later", http.StatusInternalServerError)
		return
	}

	utils.ResponseMessage(w, "Created", "Tweet created succesfully", http.StatusCreated)
}
