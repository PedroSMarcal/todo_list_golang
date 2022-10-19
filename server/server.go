package server

import (
	"log"
	"net/http"
	"os"

	"github.com/PedroSMarcal/todo_list_golang/core/todo"
	"github.com/PedroSMarcal/todo_list_golang/tools"
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

	port := os.Getenv("PORT")
	tools.SetPort(&port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func NewStart() {
	mux := http.NewServeMux()

	setRoutes(mux)

	port := os.Getenv("PORT")
	tools.SetPort(&port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
