package models

import "time"

type Tweet struct {
	UserId  string    `bson:"userId" json:"userId,omitempty"`
	Content string    `bson:"content" json:"content,omitempty"`
	Date    time.Time `bson:"date" json:"date,omitempty"`
}
