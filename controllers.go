package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/krzemLech/go-todo-app/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func addTodo(c *fiber.Ctx) error {
	var todo Todo = Todo{Completed: false}
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ "error": err.Error() })
	}
	if todo.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ "error": "Title is required" })
	}
	// ad to DB
	res, err := db.Todos.InsertOne(context.Background(), todo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{ "error": err.Error() })
	}
	todo.ID = res.InsertedID.(primitive.ObjectID).Hex()
	

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{ "data": todo, "success": true })
}

func getTodos(c *fiber.Ctx) error {
	var todos []Todo
	// get from DB
	cursor, err := db.Todos.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{ "error": err.Error() })
	}
	defer cursor.Close(context.Background()) // close cursor
	for cursor.Next(context.Background()) {
		var todo Todo
		cursor.Decode(&todo) // add error handling
		todos = append(todos, todo)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{ "data": todos, "success": true })
}

func getTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ "error": err.Error() })
	}
	var todo Todo
	// get from DB
	err = db.Todos.FindOne(context.Background(), bson.M{ "_id": objectId }).Decode(&todo)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{ "error": "Todo not found" })
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{ "data": todo, "success": true })
}

func updateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ "error": err.Error() })
	}
	filter := bson.M{ "_id": objectId }
	// TODO: should remove _id from update as an additional security measure
	update := fiber.Map{}
	if err := c.BodyParser(&update); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ "error": err.Error() })
	}
	if update["_id"] != nil {
		delete(update, "_id")
	}
	up, err := db.Todos.UpdateOne(context.Background(), filter, bson.M{ "$set": update })
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{ "error": err.Error() })
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{ "success": true, "msg": fmt.Sprintf("Modified %v document(s)", up.ModifiedCount) })
}

func deleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ "error": err.Error() })
	}
	filter := bson.M{ "_id": objectId }
	del, err := db.Todos.DeleteOne(context.Background(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{ "error": err.Error() })
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{ "success": true, "msg": fmt.Sprintf("Deleted %v document(s)", del.DeletedCount) })
}
