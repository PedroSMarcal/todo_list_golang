package server

import (
	"net/http"

	"github.com/PedroSMarcal/todo/core/todo"
)

func setRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", todo.HealthCheck)
}

func Start() {
	mux := http.NewServeMux()

	setRoutes(mux)

	http.ListenAndServe(":3333", mux)
}
