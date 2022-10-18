package main

import (
	"github.com/PedroSMarcal/todo/config"
	"github.com/PedroSMarcal/todo/database"
	"github.com/PedroSMarcal/todo/server"
)

func main() {
	config.LoadEnv()
	database.StartMigrations()
	server.Start()
}
