package todo

import (
	"io"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is my Website\n")
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Get All")
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Okay")
}
