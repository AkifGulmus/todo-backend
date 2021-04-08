package tests

import (
	"bytes"
	"encoding/json"
	router "go-backend/routes"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/stretchr/testify/assert"
)

type requestBodyModel struct {
	Text string
	Done bool
}

func TestPostRequest(t *testing.T) {

	app := fiber.New()
	app.Use(cors.New())
	router.SetupRoutes(app)

	badRequestBodyBeforeJSON := &requestBodyModel{
		Text: "",
	}
	badRequestBody, _ := json.Marshal(badRequestBodyBeforeJSON)

	goodRequestBodyBeforeJSON := &requestBodyModel{
		Text: "Buy some milk",
	}
	goodRequestBody, _ := json.Marshal(goodRequestBodyBeforeJSON)

	badRequest, err := http.NewRequest(http.MethodPost, "/todo-items", bytes.NewBuffer(badRequestBody))
	badRequest.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	badResponse, _ := app.Test(badRequest)
	assert.Equal(t, fiber.StatusBadRequest, badResponse.StatusCode)

	correctRequest, err := http.NewRequest(http.MethodPost, "/todo-items", bytes.NewBuffer(goodRequestBody))
	correctRequest.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	defer correctRequest.Body.Close()

	goodResp, _ := app.Test(correctRequest)
	assert.Equal(t, goodResp.StatusCode, fiber.StatusAccepted)
}
