package constants

import "os"

const (
	EnvError string = "Could not load the local envs"
)

var (
	AppPort string = os.Getenv("PORT")
)

var (
	TodoDatabaseDatabase string = os.Getenv("TODO_DATABASE")
	TodoDatabasePort     string = os.Getenv("TODO_PORT")
	TodoDatabaseHost     string = os.Getenv("TODO_HOST")
	TodoDatabaseUser     string = os.Getenv("TODO_USER")
	TodoDatabasePassword string = os.Getenv("TODO_PASSWORD")
)
