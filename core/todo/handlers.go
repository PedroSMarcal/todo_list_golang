package todo

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/PedroSMarcal/todo_list_golang/coll"
	"go.mongodb.org/mongo-driver/bson"
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
		task := coll.Task{}

		if id == "" {
			value, err := getAllTodosService()
			if err != nil {
				io.WriteString(w, err.Error())
				w.WriteHeader(http.StatusNoContent)
				return
			}

			b, err := bson.Marshal(value)
			if err != nil {
				io.WriteString(w, err.Error())
				return
			}

			err = bson.Unmarshal(b, &task)
			if err != nil {
				io.WriteString(w, err.Error())
				return
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

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	requestTask := RequestTask{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Bad Request")
		return
	}

	err = json.Unmarshal(body, &requestTask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Bad Request")
		return
	}

	task, err := createTodo(&requestTask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Internal Server Error")
		return
	}

	json.NewEncoder(w).Encode(&task)

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	io.WriteString(w, "Update")
}

func Finish(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	io.WriteString(w, "Finish")
}
