package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ReturnTweet struct {
	ID      primitive.ObjectID `bson:"_id" json:"idTweet,omitempty"`
	UserId  string             `bson:"userId" json:"userId,omitempty"`
	Content string             `bson:"content" json:"content,omitempty"`
	Date    time.Time          `bson:"date" json:"date,omitempty"`
}
