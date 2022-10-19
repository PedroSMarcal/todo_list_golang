package main

import (
	"github.com/PedroSMarcal/todo_list_golang/config"
	"github.com/PedroSMarcal/todo_list_golang/database"
	"github.com/PedroSMarcal/todo_list_golang/server"
)

func main() {
	config.GetEnvironmentVariables()
	database.StartMigrations()
	server.Start()
}
