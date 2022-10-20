package todo

import (
	"time"

	"github.com/PedroSMarcal/todo_list_golang/coll"
	"github.com/PedroSMarcal/todo_list_golang/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getAllTodosService() ([]coll.Task, error) {
	value, err := repository.Show()

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

	_, err := repository.PostTodo(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func getTodoById(id string, task *coll.Task) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	repository.GetByTodoId(oid, task)
	return nil
}

func updateTodoContentById(id string, content string) (*mongo.UpdateResult, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	res, err := repository.UpdateTodoById(oid, content)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func deleteTodo(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	err = repository.DeleteTodo(oid)
	if err != nil {
		return err
	}

	return nil
}
