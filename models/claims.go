package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claim struct {
	Email string             `json:"email"`
	ID    primitive.ObjectID `bson:"_ID" json:"_ID,omitempty"`
	jwt.StandardClaims
}
