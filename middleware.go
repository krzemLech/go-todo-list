package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/krzemLech/go-todo-app/config"
	"github.com/krzemLech/go-todo-app/db"
	"go.mongodb.org/mongo-driver/bson"
)

func checkMaxCount(c *fiber.Ctx) error {
	count, err := db.Todos.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{ "error": "Error checking todos" })
	}
	if int(count) >= config.Envs.MaxTodos {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{ "error": "Max todos count reached" })
	}
	return c.Next();
}