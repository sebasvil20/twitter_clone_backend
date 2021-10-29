package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GotEnvVariable(key string) string {

	// load .env file
	if os.Getenv("SCOPE") != "prod" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Error loading .env file")
		}
	}

	return os.Getenv(key)
}
