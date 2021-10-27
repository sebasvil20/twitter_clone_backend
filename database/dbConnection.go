package database

import (
	"context"
	"github.com/sebasvil20/twitter_clone_backend/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// MongoConnection connection database object
var MongoConnection = ConnectToDB()

var clientOptions = options.Client().ApplyURI("mongodb+srv://" + utils.GotEnvVariable("DB_USER") + ":" + utils.GotEnvVariable("DB_PASSWORD") + "@twittercloneapi.bmkwg.mongodb.net/" + utils.GotEnvVariable("DB_NAME") +"?retryWrites=true&w=majority")

// ConnectToDB tries to connect to db and return client or the error
func ConnectToDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf(err.Error())
		return client
	}
	log.Println("DB Connection stablished succesfully")

	return client
}

// CheckConnection tries the ping to the database
func CheckConnection() int {
	err := MongoConnection.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}