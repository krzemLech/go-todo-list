package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/krzemLech/go-todo-app/config"
	"github.com/krzemLech/go-todo-app/db"
	"go.mongodb.org/mongo-driver/mongo"
)
type Todo struct {
	ID        string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var collection *mongo.Collection

func init() {
	config.Envs.Init()
	db.ConnectMongo()
}

func main() {
	app := fiber.New()

	defer db.Client.Disconnect(context.Background())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{ "msg": "Hello, World!" })
	})

	// todos
	todoRoutes := app.Group("/todos")

	todoRoutes.Post("/", addTodo)
	todoRoutes.Get("/", getTodos)
	todoRoutes.Get("/:id", getTodo)
	todoRoutes.Patch("/:id", updateTodo);
	todoRoutes.Delete("/:id", deleteTodo)

	PORT := ":" + config.Envs.Port
	app.Listen(PORT)
}