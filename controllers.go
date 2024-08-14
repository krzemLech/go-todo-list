package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/krzemLech/go-todo-app/crud"
	"github.com/krzemLech/go-todo-app/db"
	"github.com/krzemLech/go-todo-app/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func addTodo(c *fiber.Ctx) error {
	var todo types.Todo = types.Todo{Completed: false}
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ "error": err.Error() })
	}
	if todo.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ "error": "Title is required" })
	}

	isProfane := checkProfane(todo.Title)
	if isProfane {
		return fiber.NewError(fiber.StatusBadRequest, "Profane words are not allowed")
	}
	// ad to DB
	res, err := crud.AddTodo(todo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{ "error": err.Error() })
	}
	todo.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{ "data": todo, "success": true })
}

func getTodos(c *fiber.Ctx) error {
	pageParam := c.Query("page", "1")
	perPageParam := c.Query("perPage", "5")
	page, perPage, err := formatFilters(pageParam, perPageParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ "error": err.Error() })
	}

	// get from DB
	var todos []types.Todo
	count, err := crud.CountTodos()
	cursor, err2 := crud.GetTodos(page, perPage)
	if err != nil || err2 != nil {
		fmt.Println("counting", err)
		fmt.Println("finding", err2)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{ "error": "Error fetching todos" })
	}
	defer cursor.Close(context.Background()) // close cursor

	// format data
	for cursor.Next(context.Background()) {
		var todo types.Todo
		cursor.Decode(&todo) // add error handling
		todos = append(todos, todo)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{ "data": todos, "success": true, "total": count })
}

func getTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ "error": err.Error() })
	}
	var todo types.Todo
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
	update := fiber.Map{}
	if err := c.BodyParser(&update); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ "error": err.Error() })
	}
	if update["_id"] != nil {
		delete(update, "_id")
	}
	up, err := crud.UpdateTodo(id, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{ "error": err.Error() })
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{ "success": true, "msg": fmt.Sprintf("Modified %v document(s)", up.ModifiedCount) })
}

func deleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	del, err := crud.DeleteTodo(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{ "error": err.Error() })
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{ "success": true, "msg": fmt.Sprintf("Deleted %v document(s)", del.DeletedCount) })
}
