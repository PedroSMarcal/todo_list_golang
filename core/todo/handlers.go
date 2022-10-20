package todo

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/PedroSMarcal/todo_list_golang/coll"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

	strings.Split(r.RequestURI, "/")
	io.WriteString(w, "This is my Website\n ")
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

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

			json.NewEncoder(w).Encode(&value)
			return
		}

		err := getTodoById(id, &task)
		if err != nil {
			io.WriteString(w, err.Error())
			w.WriteHeader(http.StatusNoContent)
			return
		}

		json.NewEncoder(w).Encode(&task)

	case "OPTIONS":
		w.WriteHeader(http.StatusAccepted)
		io.WriteString(w, "Bem Vindo")
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, "Not Allowed")
		return
	}

}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

	switch r.Method {
	case "POST":
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

	case "OPTIONS":
		w.WriteHeader(http.StatusAccepted)
		io.WriteString(w, "Bem Vindo")
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, "Not Allowed")
		return
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

	switch r.Method {
	case "PUT":
		values := r.URL.Query()
		id := values.Get("id")

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

		res, err := updateTodoContentById(id, requestTask.Content)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "Bad Request")
			return
		}

		json.NewEncoder(w).Encode(&res)

	case "OPTIONS":
		w.WriteHeader(http.StatusAccepted)
		io.WriteString(w, "Bem Vindo")
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, "Not Allowed")
		return
	}
}

func Finish(w http.ResponseWriter, r *http.Request) {
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

	switch r.Method {
	case "DELETE":
		values := r.URL.Query()
		id := values.Get("id")

		err := deleteTodo(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "Bad Request")
			return
		}

		io.WriteString(w, "FInish with Success")

	case "OPTIONS":
		w.WriteHeader(http.StatusAccepted)
		io.WriteString(w, "Bem Vindo")
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, "Not Allowed")
		return
	}
}
