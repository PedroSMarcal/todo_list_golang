package todo

import (
	"github.com/PedroSMarcal/todo/models"
	"github.com/PedroSMarcal/todo/repository"
)

var repo repository.TodoRepository

func getAllTodosService(tasks *models.Task) error {
	err := repo.GetAll(tasks)

	if err != nil {
		return err
	}

	return nil
}
