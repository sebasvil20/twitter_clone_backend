package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
