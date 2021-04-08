package tests

import (
	router "go-backend/routes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/stretchr/testify/assert"
)

func TestGetRequest(t *testing.T) {

	app := fiber.New()
	app.Use(cors.New())
	router.SetupRoutes(app)

	req, err := http.NewRequest(http.MethodGet, "/todo-items", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, _ := app.Test(req)

	var expectedResponse []byte

	assert.Equal(t, resp.StatusCode, fiber.StatusOK)
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, reflect.TypeOf(body), reflect.TypeOf(expectedResponse))

}
