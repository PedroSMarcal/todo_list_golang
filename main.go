package main

import (
	"fmt"

	"github.com/PedroSMarcal/todo/config"
	"github.com/PedroSMarcal/todo/database"
	"github.com/PedroSMarcal/todo/server"
)

func main() {
	config.LoadEnv()
	database.StartMigrations()
	fmt.Print("HI")
	server.Start()
}
