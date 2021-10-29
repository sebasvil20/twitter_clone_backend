package repository

import (
	"context"
	"github.com/sebasvil20/twitter_clone_backend/database"
	"github.com/sebasvil20/twitter_clone_backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func SaveNewTweet(tweet models.Tweet) (string, bool, error){
	rContext, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := database.MongoConnection.Database("twitterclone")
	collection := db.Collection("tweets")

	document := bson.M{
		"userId" : tweet.UserId,
		"content" : tweet.Content,
		"date": tweet.Date,
	}

	result, err := collection.InsertOne(rContext, document)
	if err != nil {
		return "", false, err
	}

	objId, _ := result.InsertedID.(primitive.ObjectID)
	return objId.String(), true, nil
}