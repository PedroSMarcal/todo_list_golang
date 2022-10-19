package server

import (
	"log"
	"net/http"
	"os"

	"github.com/PedroSMarcal/todo/core/todo"
)

func setRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/create", todo.CreateTask)
	mux.HandleFunc("/update", todo.UpdateTask)
	mux.HandleFunc("/finish", todo.Finish)
	mux.HandleFunc("/get", todo.GetTask)

	mux.HandleFunc("/health", todo.HealthCheck)
}

func Start() {
	mux := http.NewServeMux()

	setRoutes(mux)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), mux)
	if err != nil {
		log.Fatal(err)
	}
}
