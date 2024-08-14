package crud

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/krzemLech/go-todo-app/db"
	"github.com/krzemLech/go-todo-app/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CountTodos() (int64, error) {
	// Count all todos
	return db.Todos.CountDocuments(context.Background(), bson.M{})
}

func AddTodo(todo types.Todo) (*mongo.InsertOneResult, error) {
	// Add a todo
	return db.Todos.InsertOne(context.Background(), todo)
}

func GetTodos(page, perPage int64) (*mongo.Cursor, error) {
	opts := options.Find().SetSkip((page - 1) * perPage).SetLimit(perPage)
	return db.Todos.Find(context.Background(), bson.M{}, opts)
}

func UpdateTodo(id string, update fiber.Map) (*mongo.UpdateResult, error) {
	formattedID, err := db.ConvertToID(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{ "_id": formattedID }
	// Update a todo
	return db.Todos.UpdateOne(context.Background(), filter, bson.M{ "$set": update })
}

func DeleteTodo(id string) (*mongo.DeleteResult, error) {
	formattedID, err := db.ConvertToID(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{ "_id": formattedID }
	// Delete a todo
	return db.Todos.DeleteOne(context.Background(), filter)
}