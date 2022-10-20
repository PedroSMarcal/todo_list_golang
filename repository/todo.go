package repository

import (
	"context"

	"github.com/PedroSMarcal/todo_list_golang/coll"
	"github.com/PedroSMarcal/todo_list_golang/config"
	"github.com/PedroSMarcal/todo_list_golang/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// type (
// 	TodoRepository interface {
// 		ShowRepository() ([]primitive.M, error)
// 		// PostRepository(task *coll.Task) (*mongo.InsertOneResult, error)
// 		// GetAnyReposity(key, value string) (*mongo.SingleResult, error)
// 		// UpdateRepository(element coll.Task) (*mongo.UpdateResult, error)
// 	}

// 	todoRepository struct {
// 		db *mongo.Database
// 	}
// )

func getCollection(r *mongo.Database) *mongo.Collection {
	return r.Collection(config.EnvVariable.CollectionName)
}

func Show() ([]coll.Task, error) {
	connection := database.ConnectDatabase()
	collection := getCollection(connection)

	filter := bson.M{}
	result := []coll.Task{}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	if err := cur.All(context.TODO(), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func PostTodo(task *coll.Task) (*mongo.InsertOneResult, error) {
	connection := database.ConnectDatabase()
	collection := getCollection(connection)

	res, err := collection.InsertOne(context.TODO(), &task)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetByTodoId(oid primitive.ObjectID, task *coll.Task) {
	connection := database.ConnectDatabase()
	collection := getCollection(connection)

	filter := bson.M{"_id": oid}

	collection.FindOne(context.TODO(), filter).Decode(&task)
}

func UpdateTodoById(oid primitive.ObjectID, content string) (*mongo.UpdateResult, error) {
	connection := database.ConnectDatabase()
	collection := getCollection(connection)

	filter := bson.M{"_id": oid}
	fields := bson.M{"$set": bson.M{"content": content}}

	res, err := collection.UpdateOne(context.TODO(), filter, fields)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func DeleteTodo(oid primitive.ObjectID) error {
	connection := database.ConnectDatabase()
	collection := getCollection(connection)

	filter := bson.M{"_id": oid}
	field := bson.M{"$set": bson.M{"finish": true}}

	_, err := collection.UpdateOne(context.TODO(), filter, field)
	if err != nil {
		return err
	}

	return nil
}
