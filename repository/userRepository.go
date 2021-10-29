package repository

import (
	"context"
	"github.com/sebasvil20/twitter_clone_backend/database"
	"github.com/sebasvil20/twitter_clone_backend/models"
	"github.com/sebasvil20/twitter_clone_backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func SaveNewUser(user models.User) (bool, error) {

	rContext, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := database.MongoConnection.Database("twitterclone")
	collection := db.Collection("users")

	user.Password, _ = utils.EncryptPassword(user.Password)

	_, err := collection.InsertOne(rContext, user)
	if err != nil {
		return false, err
	}

	return true, nil
}

func FindUserByEmail(email string) (models.User, bool, error) {
	rContext, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := database.MongoConnection.Database("twitterclone")
	collection := db.Collection("users")

	criteria := bson.M{
		"email": email,
	}

	var result models.User

	err := collection.FindOne(rContext, criteria).Decode(&result)
	if err != nil {
		return result, false, err
	}

	return result, true, nil
}

func FindProfileByID(Id string) (models.User, error) {
	rContext, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := database.MongoConnection.Database("twitterclone")
	collection := db.Collection("users")

	var profile models.User

	objID, _ := primitive.ObjectIDFromHex(Id)

	criteria := bson.M{
		"_id": objID,
	}

	err := collection.FindOne(rContext, criteria).Decode(&profile)
	profile.Password = ""
	if err != nil {
		return profile, err
	}

	return profile, nil

}
