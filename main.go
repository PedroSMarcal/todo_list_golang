package main

import (
	"github.com/PedroSMarcal/todo/database"
	"github.com/PedroSMarcal/todo/server"
)

func main() {
	database.StartMigrations()
	server.Start()
}
