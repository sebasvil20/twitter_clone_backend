package repository

import (
	"context"
	"github.com/sebasvil20/twitter_clone_backend/database"
	"github.com/sebasvil20/twitter_clone_backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func SaveNewTweet(tweet models.Tweet) (string, bool, error) {
	rContext, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := database.MongoConnection.Database("twitterclone")
	collection := db.Collection("tweets")

	document := bson.M{
		"userId":  tweet.UserId,
		"content": tweet.Content,
		"date":    tweet.Date,
	}

	result, err := collection.InsertOne(rContext, document)
	if err != nil {
		return "", false, err
	}

	objId, _ := result.InsertedID.(primitive.ObjectID)
	return objId.String(), true, nil
}

func GetTweetsByUserId(idUser string, page int64) ([]*models.ReturnTweet, bool) {
	rContext, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := database.MongoConnection.Database("twitterclone")
	collection := db.Collection("tweets")

	var result []*models.ReturnTweet

	var criteria = bson.M{
		"userId": idUser,
	}

	queryOptions := options.Find()
	queryOptions.SetLimit(20).SetSort(bson.D{{Key: "date", Value: -1}}).SetSkip((page - 1) * 20)

	pointer, err := collection.Find(rContext, criteria, queryOptions)
	if err != nil {
		return result, false
	}

	for pointer.Next(context.TODO()) {
		var document models.ReturnTweet
		err := pointer.Decode(&document)
		if err != nil {
			return result, false
		}
		result = append(result, &document)
	}

	return result, true
}
