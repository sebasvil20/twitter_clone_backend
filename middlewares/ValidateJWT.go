package middlewares

import (
	"github.com/sebasvil20/twitter_clone_backend/auth"
	"github.com/sebasvil20/twitter_clone_backend/utils"
	"net/http"
)

func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _, _, err := auth.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			utils.ResponseMessage(w, "InvalidToken", err.Error()+", Enter a valid JWT and try again", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
