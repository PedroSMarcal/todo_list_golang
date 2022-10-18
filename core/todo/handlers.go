package todo

import (
	"io"
	"net/http"
	"strings"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	strings.Split(r.RequestURI, "/")
	io.WriteString(w, "This is my Website\n ")
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values.Get("id")

	if id == "" {
		io.WriteString(w, "Get All")
		return
	}

	io.WriteString(w, "Get "+id)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Created")
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Update")
}

func Finish(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Finish")
}
