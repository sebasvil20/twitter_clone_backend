package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sebasvil20/twitter_clone_backend/models"
	"github.com/sebasvil20/twitter_clone_backend/utils"
	"log"
	"time"
)

func GenerateJWT(user models.User) (string, error) {
	secretWord := []byte(utils.GotEnvVariable("JWT_SECRET"))
	log.Println(string(secretWord))

	payload := jwt.MapClaims{
		"email":       user.Email,
		"firstName":   user.FirstName,
		"lastName":    user.LastName,
		"dateOfBirth": user.DateOfBirth,
		"biography":   user.Biography,
		"location":    user.Location,
		"website":     user.Website,
		"_id":         user.ID.Hex(),
		"exp":         time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, _ := token.SignedString(secretWord)

	return tokenStr, nil
}
