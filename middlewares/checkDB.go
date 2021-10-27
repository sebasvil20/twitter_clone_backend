package middlewares

import (
	"github.com/sebasvil20/twitter_clone_backend/database"
	"net/http"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if database.CheckConnection() == 0 {
			http.Error(w, "Can't connect to database, try again later", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	}
}
