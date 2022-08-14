package app

import (
	"ledungcobra/gateway-go/pkg/config"
	"ledungcobra/gateway-go/pkg/middlewares"
	"ledungcobra/gateway-go/pkg/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type IServer interface {
	Initialize() error
	Listen() error
	Stop() error
}

func NewServer() IServer {
	return &App{}
}

type App struct {
	port   string
	server *fiber.App
}

// Initialize
func (a *App) Initialize() error {
	log.Println("Initializing server...")
	log.Println("Loading env from .env file")
	err := godotenv.Load(".env")

	if err != nil {
		log.Print("Error loading .env file", err)
	}
	config.Load()
	config := config.Config
	a.port = config.Port
	a.server = fiber.New()
	routes.SetUpRoutes(a.server)
	middlewares.SetUpMiddlewares(a.server)
	return nil
}

// Listen
func (a *App) Listen() error {
	log.Print("Listening on port " + a.port + "...")
	return a.server.Listen(a.getConnection())
}

func (a *App) getConnection() string {
	return ":" + a.port
}

// Stop
func (*App) Stop() error {
	return nil
}
