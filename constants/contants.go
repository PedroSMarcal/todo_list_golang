package constants

import "os"

const (
	EnvError string = "Could not load the local envs"
)

var (
	AppPort string = os.Getenv("PORT")
)
