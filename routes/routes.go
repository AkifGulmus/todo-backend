package router

import (
	"go-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	type Request struct {
		Text string
		Done bool
	}

	app.Get("/todo-items", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Status(fiber.StatusOK).JSON(services.GetTodos())
	})

	app.Post("/todo-items", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		body := new(Request)
		if err := c.BodyParser(body); err != nil {
			return err
		}

		if body.Text == "" {
			return c.Status(fiber.StatusBadRequest).Send(nil)
		}
		for key := range services.Todos {
			if services.Todos[key].Text == body.Text {
				return c.Status(fiber.StatusBadRequest).Send(nil)
			}
		}
		services.CreateTodo(body.Text)
		return c.Status(fiber.StatusAccepted).Send(nil)
	})

	app.Patch("/todo-items/:todoid", func(c *fiber.Ctx) error {
		if _, ok := services.Todos[c.Params("todoid")]; !ok {
			return c.Status(fiber.StatusBadRequest).Send(nil)
		}
		c.Set("Content-Type", "application/json")
		services.UpdateTodo(c.Params("todoid"))
		return c.Status(fiber.StatusAccepted).Send(nil)
	})

	app.Delete("/todo-items/:todoid", func(c *fiber.Ctx) error {
		if _, ok := services.Todos[c.Params("todoid")]; !ok {
			return c.Status(fiber.StatusBadRequest).Send(nil)
		}
		c.Set("Content-Type", "application/json")
		services.DeleteTodo(c.Params("todoid"))
		return c.Status(fiber.StatusAccepted).Send(nil)
	})

}
