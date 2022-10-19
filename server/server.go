package server

import (
	"log"
	"net/http"

	"github.com/PedroSMarcal/todo/config"
	"github.com/PedroSMarcal/todo/core/todo"
	"github.com/PedroSMarcal/todo/tools"
)

func setRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/create", todo.CreateTask)
	mux.HandleFunc("/update", todo.UpdateTask)
	mux.HandleFunc("/finish", todo.Finish)
	mux.HandleFunc("/", todo.GetTask)

	mux.HandleFunc("/health", todo.HealthCheck)
}

func Start() {
	mux := http.NewServeMux()

	setRoutes(mux)

	port := config.Envs.PORT
	tools.SetPort(&port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
