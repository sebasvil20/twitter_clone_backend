package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/sebasvil20/twitter_clone_backend/models"
	"github.com/sebasvil20/twitter_clone_backend/repository"
	"github.com/sebasvil20/twitter_clone_backend/utils"
	"strings"
)

/*Email valor de Email usado en todos los EndPoints */
var Email string

/*IDUsuario es el ID devuelto del modelo, que se usará en todos los EndPoints */
var IdUser string

/*ProcesoToken proceso token para extraer sus valores */
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	secretWord := []byte(utils.GotEnvVariable("JWT_SECRET"))
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("invalid token format or token is empty")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return secretWord, nil
	})
	if err == nil {
		_, found, _ := repository.FindUserByEmail(claims.Email)
		if found == true {
			Email = claims.Email
			IdUser = claims.ID.Hex()
		}
		return claims, found, IdUser, nil
	}
	if !tkn.Valid {
		return claims, false, "", errors.New("invalid token")
	}
	return claims, false, "", err
}
