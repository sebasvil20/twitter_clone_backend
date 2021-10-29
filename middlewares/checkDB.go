package middlewares

import (
	"github.com/sebasvil20/twitter_clone_backend/database"
	"github.com/sebasvil20/twitter_clone_backend/utils"
	"net/http"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if database.CheckConnection() == 0 {
			utils.ResponseMessage(w, "InternalError", "Can't connect to database, try again later", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
