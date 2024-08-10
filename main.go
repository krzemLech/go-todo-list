package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/krzemLech/go-todo-app/config"
)
type Todo struct {
	ID        string    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func init() {
	config.Envs.Init()
}

func main() {
	app := fiber.New()

	todos := []Todo{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{ "msg": "Hello, World!" })
	})

	// todos
	todoRoutes := app.Group("/todos")

	todoRoutes.Post("/", func(c *fiber.Ctx) error {
		// id := fmt.Sprintf("%v", rand.Intn(100000000))
		var todo Todo = Todo{Completed: false, ID: fmt.Sprint(rand.Intn(100000000)), Title: ""}
		if err := c.BodyParser(&todo); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ "error": err.Error() })
		}
		if todo.Title == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ "error": "Title is required" })
		}
		todos = append(todos, todo)

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{ "data": todo, "success": true })
	})
	todoRoutes.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{ "data": todos, "success": true })
	})
	todoRoutes.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		log.Println(id)
		log.Println(todos[0].ID)
		log.Println(id == todos[0].ID)
		for _, todo := range todos {
			if id == string(todo.ID) {
				return c.Status(fiber.StatusOK).JSON(fiber.Map{ "data": todo, "success": true })
			}
		}

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{ "error": "Todo not found" })
	})

	todoRoutes.Patch("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var newTodo Todo
		if err := c.BodyParser(&newTodo); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ "error": err.Error() })
		}
		for i, todo := range todos {
			if id == string(todo.ID) {
				todos[i].Title = newTodo.Title
				todos[i].Completed = newTodo.Completed
				return c.Status(fiber.StatusOK).JSON(fiber.Map{ "data": todos[i], "success": true })
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{ "error": "Todo not found" })
	});

	todoRoutes.Delete("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, todo := range todos {
			if id == string(todo.ID) {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(fiber.StatusOK).JSON(fiber.Map{ "success": true })
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{ "error": "Todo not found" })
	})

	PORT := ":" + config.Envs.Port
	app.Listen(PORT)
}