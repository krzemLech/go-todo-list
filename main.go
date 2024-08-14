package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/krzemLech/go-todo-app/config"
	"github.com/krzemLech/go-todo-app/db"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func init() {
	config.Envs.Init()
	db.ConnectMongo()
}

func main() {
	app := fiber.New()

	if config.Envs.Env == "development" {
		app.Use(cors.New())
	}

	defer db.Client.Disconnect(context.Background())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{ "msg": "works" })
	})

	// todos
	todoRoutes := app.Group("/api/v1/todos")

	todoRoutes.Post("/", checkMaxCount, addTodo)
	todoRoutes.Get("/", getTodos)
	todoRoutes.Get("/:id", getTodo)
	todoRoutes.Patch("/:id", updateTodo);
	todoRoutes.Delete("/:id", deleteTodo)

	if config.Envs.Env == "production" {
		app.Static("/", "./client/dist")
	}

	PORT := ":" + config.Envs.Port
	app.Listen(PORT)
}