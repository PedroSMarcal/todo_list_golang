package todo

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/PedroSMarcal/todo_list_golang/coll"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	strings.Split(r.RequestURI, "/")
	io.WriteString(w, "This is my Website\n ")
}

func GetTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	switch r.Method {
	case "GET":
		values := r.URL.Query()
		id := values.Get("id")

		if id == "" {
			value, err := getAllTodosService()
			if err != nil {
				io.WriteString(w, err.Error())
				w.WriteHeader(http.StatusNoContent)
				return
			}

			task := coll.Task{}
			s, err := json.Marshal(value)
			if err != nil {
				log.Print(err)
			}

			err = json.Unmarshal(s, &task)
			if err != nil {
				log.Print(err)
			}

			json.NewEncoder(w).Encode(&task)
			return
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
