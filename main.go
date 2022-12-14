package main

import (
	"github.com/PedroSMarcal/todo_list_golang/config"
	"github.com/PedroSMarcal/todo_list_golang/server"
)

func main() {
	config.GetEnvironmentVariables()
	server.Start()
}
