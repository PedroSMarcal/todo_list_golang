package config

import (
	"log"

	"github.com/PedroSMarcal/todo/constants"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(constants.EnvError)
	}
}
