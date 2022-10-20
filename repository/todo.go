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

func ShowRepository() ([]primitive.M, error) {
	connection := database.ConnectDatabase()
	collection := getCollection(connection)

	filter := bson.M{}
	result := []bson.M{}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	if err := cur.All(context.Background(), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func PostRepository(task *coll.Task) (*mongo.InsertOneResult, error) {
	connection := database.ConnectDatabase()
	collection := getCollection(connection)

	res, err := collection.InsertOne(context.Background(), &task)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// func (r *todoRepository) GetAnyReposity(key, value string) (*mongo.SingleResult, error) {
// 	collection := r.getCollection()
// 	filter := bson.M{key: value}

// 	coll := collection.FindOne(context.TODO(), filter)

// 	return coll, nil
// }

// func (r *todoRepository) UpdateRepository(element coll.Task) (*mongo.UpdateResult, error) {
// 	collection := r.getCollection()
// 	filter := bson.M{element.ID: element.ID}
// 	fields := bson.M{"$set": element}

// 	res, err := collection.UpdateOne(context.Background(), filter, fields)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return res, nil
// }

// func (r *todoRepository) DeleteRepository(key, value string) (*mongo.DeleteResult, error) {
// 	collection := r.getCollection()
// 	filter := bson.M{key: value}

// 	res, err := collection.DeleteOne(context.Background(), filter)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return res, nil
// }
