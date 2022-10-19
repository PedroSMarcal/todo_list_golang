package config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentVar struct {
	MongoURI       string
	Port           string
	DatabaseName   string
	CollectionName string
}

var EnvVariable *EnvironmentVar

func GetEnvironmentVariables() *EnvironmentVar {
	godotenv.Load(".env")

	EnvVariable = &EnvironmentVar{
		MongoURI:       os.Getenv("MONGO_URI"),
		DatabaseName:   os.Getenv("DATABASE_NAME"),
		CollectionName: os.Getenv("COLLECTION_NAME"),
		Port:           os.Getenv("PORT"),
	}

	return EnvVariable
}
