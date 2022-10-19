package config

import (
	"log"
	"os"

	"github.com/PedroSMarcal/todo/constants"
	"github.com/joho/godotenv"
)

var Envs EnvsStruct

type EnvsStruct struct {
	PORTS string
	DSN   string
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(constants.EnvError)
	}

	Envs.DSN = os.Getenv("DATABASE_URL")
	Envs.PORTS = os.Getenv("PORT")
}
