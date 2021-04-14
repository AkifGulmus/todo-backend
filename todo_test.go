package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/stretchr/testify/assert"
)

func TestPostRequest(t *testing.T) {

	app := fiber.New()
	app.Use(cors.New())
	SetupRoutes(app)

	badRequestBodyBeforeJSON := &requestBodyModel{
		Text: "",
	}
	badRequestBody, _ := json.Marshal(badRequestBodyBeforeJSON)

	badRequest, err := http.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(badRequestBody))
	badRequest.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	badResponse, _ := app.Test(badRequest)
	assert.Equal(t, fiber.StatusBadRequest, badResponse.StatusCode)

	goodRequestBodyBeforeJSON := &requestBodyModel{
		Text: "Buy some milk",
	}
	goodRequestBody, _ := json.Marshal(goodRequestBodyBeforeJSON)

	goodRequest, err := http.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(goodRequestBody))
	goodRequest.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	defer goodRequest.Body.Close()

	goodResp, _ := app.Test(goodRequest)
	assert.Equal(t, goodResp.StatusCode, fiber.StatusCreated)
}

func TestDeleteRequest(t *testing.T) {
	app := fiber.New()
	app.Use(cors.New())
	SetupRoutes(app)

	badRequest, err := http.NewRequest(http.MethodDelete, "/todos/sometodoid", nil)
	badRequest.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	badResp, _ := app.Test(badRequest)

	assert.Equal(t, fiber.StatusBadRequest, badResp.StatusCode)

	Todos["sometodoid"] = Todo{
		TodoID: "sometodoid",
		Text:   "Buy some milk",
		Done:   false,
	}

	goodRequest, err := http.NewRequest(http.MethodDelete, "/todos/sometodoid", nil)
	goodRequest.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	goodResp, _ := app.Test(goodRequest)

	assert.Equal(t, fiber.StatusAccepted, goodResp.StatusCode)
}

func TestGetRequest(t *testing.T) {
	app := fiber.New()
	app.Use(cors.New())
	SetupRoutes(app)

	req, err := http.NewRequest(http.MethodGet, "/todos", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, _ := app.Test(req)

	var expectedResponse []byte

	assert.Equal(t, resp.StatusCode, fiber.StatusOK)
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, reflect.TypeOf(body), reflect.TypeOf(expectedResponse))
}

func TestPutRequest(t *testing.T) {
	app := fiber.New()
	app.Use(cors.New())
	SetupRoutes(app)

	badRequest, err := http.NewRequest(http.MethodPatch, "/todos/sometodoid", nil)
	badRequest.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	badResp, _ := app.Test(badRequest)

	assert.Equal(t, fiber.StatusBadRequest, badResp.StatusCode)

	Todos["sometodoid"] = Todo{
		TodoID: "sometodoid",
		Text:   "Buy some milk",
		Done:   false,
	}

	goodRequest, err := http.NewRequest(http.MethodPatch, "/todos/sometodoid", nil)
	goodRequest.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	goodResp, _ := app.Test(goodRequest)

	assert.Equal(t, fiber.StatusAccepted, goodResp.StatusCode)
}
