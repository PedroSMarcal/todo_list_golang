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

	switch r.Method {
	case "GET":
		values := r.URL.Query()
		id := values.Get("id")
		// tasks := models.Task{}

		if id == "" {
			// if err := getAllTodosService(&tasks); err != nil {
			w.WriteHeader(http.StatusNoContent)
			io.WriteString(w, "Nothing found")
			return
			// }

			// w.WriteHeader(http.StatusCreated)
			// json.NewEncoder(w).Encode(tasks)
			// return
		}

		io.WriteString(w, "Get "+id)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, "Not Allowed")
		return
	}

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
