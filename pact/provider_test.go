package pact

import (
	"fmt"
	router "go-backend/routes"
	"go-backend/services"
	"log"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

var dir, _ = os.Getwd()
var pactDir = fmt.Sprintf("%s/../pacts", dir)
var logDir = fmt.Sprintf("%s/log", dir)

var stateHandlers = types.StateHandlers{
	"A list of todos with one todo item": func() error {
		var TodosMockOneItem = make(map[string]services.Todo)
		TodosMockOneItem["0f944e2d-77d4-4063-ac35-62342de49ad2"] = services.Todo{
			TodoID: "0f944e2d-77d4-4063-ac35-62342de49ad2",
			Text:   "Buy some milk",
			Done:   false,
		}
		services.Todos = TodosMockOneItem
		return nil
	},
	"A new todo item is added": func() error {
		var EmptyTodosMock = make(map[string]services.Todo)
		services.Todos = EmptyTodosMock
		return nil
	},
	"There is a todo item to update the status of": func() error {
		var TodosMockOneItem = make(map[string]services.Todo)
		TodosMockOneItem["0f944e2d-77d4-4063-ac35-62342de49ad2"] = services.Todo{
			TodoID: "0f944e2d-77d4-4063-ac35-62342de49ad2",
			Text:   "Buy some milk",
			Done:   false,
		}
		services.Todos = TodosMockOneItem
		return nil
	},
	"A todo item to be deleted": func() error {
		var TodosMockOneItem = make(map[string]services.Todo)
		TodosMockOneItem["0f944e2d-77d4-4063-ac35-62342de49ad2"] = services.Todo{
			TodoID: "0f944e2d-77d4-4063-ac35-62342de49ad2",
			Text:   "Buy some milk",
			Done:   false,
		}
		services.Todos = TodosMockOneItem
		return nil
	},
}

func TestPactProvider(t *testing.T) {
	go startProviderServer()
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	pact := dsl.Pact{
		Provider:                 "backend",
		LogDir:                   logDir,
		PactDir:                  pactDir,
		DisableToolValidityCheck: true,
		LogLevel:                 "INFO",
	}

	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:            "http://localhost:5000",
		FailIfNoPactsFound:         true,
		StateHandlers:              stateHandlers,
		BrokerURL:                  os.Getenv("PACT_BROKER_URL"),
		BrokerToken:                os.Getenv("PACT_BROKER_TOKEN"),
		PublishVerificationResults: true,
		ProviderVersion:            "1.0.0",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func startProviderServer() {
	log.Printf("API starting: port 5000")

	app := fiber.New()
	app.Use(cors.New())
	router.SetupRoutes(app)

	log.Fatal(app.Listen(":5000"))
	defer app.Shutdown()
	log.Printf("API terminating")
}
