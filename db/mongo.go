package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Todos *mongo.Collection

func ConnectMongo() {
	var err error
	mongoUri := os.Getenv("MONGO_URI")
	fmt.Println(mongoUri)
	clientOptions := options.Client().ApplyURI(mongoUri)
	Client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err2 := Client.Ping(context.Background(), nil)
	if err2 != nil {
		log.Fatal(err2)
	}
	Todos = Client.Database("go_test").Collection("todos")
	fmt.Println("Connected to MongoDB!")
}