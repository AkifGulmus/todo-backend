package main

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	type Request struct {
		Text string
		Done bool
	}
	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(GetTodos())
	})

	app.Post("/todos", func(c *fiber.Ctx) error {
		body := new(Request)
		if err := c.BodyParser(body); err != nil {

			return c.Status(fiber.StatusBadRequest).Send(nil)
		}
		if body.Text == "" {
			return c.Status(fiber.StatusBadRequest).Send(nil)
		}
		CreateTodo(body.Text)
		return c.Status(fiber.StatusCreated).Send(nil)
	})

	app.Patch("/todos/:todoid", func(c *fiber.Ctx) error {
		if _, ok := Todos[c.Params("todoid")]; !ok {
			return c.Status(fiber.StatusBadRequest).Send(nil)
		}
		UpdateTodo(c.Params("todoid"))
		return c.Status(fiber.StatusAccepted).Send(nil)
	})

	app.Delete("/todos/:todoid", func(c *fiber.Ctx) error {
		if _, ok := Todos[c.Params("todoid")]; !ok {
			return c.Status(fiber.StatusBadRequest).Send(nil)
		}
		DeleteTodo(c.Params("todoid"))
		return c.Status(fiber.StatusAccepted).Send(nil)
	})

}
