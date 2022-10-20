package todo

import (
	"time"

	"github.com/PedroSMarcal/todo_list_golang/coll"
	"github.com/PedroSMarcal/todo_list_golang/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getAllTodosService() ([]primitive.M, error) {
	value, err := repository.ShowRepository()

	if err != nil {
		return nil, err
	}

	return value, nil
}

func createTodo(req *RequestTask) (*coll.Task, error) {

	task := coll.Task{
		ID:        primitive.NewObjectID(),
		Content:   req.Content,
		Finish:    false,
		CreatedAt: time.Now(),
	}

	_, err := repository.PostRepository(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}
