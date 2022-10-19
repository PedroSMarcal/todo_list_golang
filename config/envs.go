package config

// import (
// 	"log"
// 	"os"

// 	"github.com/PedroSMarcal/todo_list_golang/constants"
// 	"github.com/joho/godotenv"
// )

// var Envs AppEnvs
// var TodoDatabase TodoDatabaseEnvs

// type TodoDatabaseEnvs struct {
// 	User     string
// 	Password string
// 	Host     string
// 	Port     string
// 	Database string
// }

// type AppEnvs struct {
// 	PORT string
// }

// func LoadEnv() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal(constants.EnvError)
// 	}

// 	TodoDatabase.User = os.Getenv("TODO_USER")
// 	TodoDatabase.Password = os.Getenv("TODO_PASSWORD")
// 	TodoDatabase.Host = os.Getenv("TODO_HOST")
// 	TodoDatabase.Port = os.Getenv("TODO_PORT")
// 	TodoDatabase.Database = os.Getenv("TODO_DATABASE")

// 	Envs.PORT = os.Getenv("PORT")
// }
