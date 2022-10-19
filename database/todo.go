package database

import (
	"context"
	"log"

	"github.com/PedroSMarcal/todo_list_golang/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoDB *mongo.Database

func MongoClient() *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.EnvVariable.MongoURI))
	if err != nil {
		log.Fatal("Invalid connection")
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
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
