package tests

import (
	router "go-backend/routes"
	"go-backend/services"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/stretchr/testify/assert"
)

func TestDeleteRequest(t *testing.T) {

	app := fiber.New()
	app.Use(cors.New())
	router.SetupRoutes(app)

	badRequest, err := http.NewRequest(http.MethodDelete, "/todo-items/sometodoid", nil)
	badRequest.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	badResp, _ := app.Test(badRequest)

	assert.Equal(t, fiber.StatusBadRequest, badResp.StatusCode)

	services.Todos["sometodoid"] = services.Todo{
		TodoID: "sometodoid",
		Text:   "Buy some milk",
		Done:   false,
	}

	correctRequest, err := http.NewRequest(http.MethodDelete, "/todo-items/sometodoid", nil)
	correctRequest.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	goodResp, _ := app.Test(correctRequest)

	assert.Equal(t, fiber.StatusAccepted, goodResp.StatusCode)
}
