package todo

import (
	"github.com/PedroSMarcal/todo_list_golang/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// var repo repository.TodoRepository

func getAllTodosService() ([]primitive.M, error) {
	value, err := repository.ShowRepository()

	if err != nil {
		return nil, err
	}

	return value, nil
}
