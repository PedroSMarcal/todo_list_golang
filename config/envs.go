package config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentVar struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
	URI        string
	Port       string
}

var EnvVariable *EnvironmentVar

func GetEnvironmentVariables() *EnvironmentVar {
	godotenv.Load(".env")

	EnvVariable = &EnvironmentVar{
		DbHost:     os.Getenv("TODO_HOST"),
		DbName:     os.Getenv("TODO_DATABASE"),
		DbPort:     os.Getenv("TODO_PORT"),
		DbUser:     os.Getenv("TODO_USER"),
		DbPassword: os.Getenv("TODO_PASSWORD"),
		URI:        os.Getenv("DATABASE_URL"),
		Port:       os.Getenv("PORT"),
	}

	return EnvVariable
}
