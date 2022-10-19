package database

import (
	"context"
	"log"

	"github.com/PedroSMarcal/todo_list_golang/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDB *mongo.Database

func MongoClient() *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.EnvVariable.MongoURI))
	if err != nil {
		log.Fatal("Invalid connection")
	}

	mongoDB = client.Database(config.EnvVariable.DatabaseName)

	return mongoDB
}

func ConnectDatabase() *mongo.Database {
	if mongoDB == nil {
		MongoClient()
	}
	return mongoDB
}
