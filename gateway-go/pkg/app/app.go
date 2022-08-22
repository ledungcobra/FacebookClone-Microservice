package app

import (
	"ledungcobra/gateway-go/pkg/app_service"
	"ledungcobra/gateway-go/pkg/interfaces"
	"log"

	"ledungcobra/gateway-go/pkg/config"
	"ledungcobra/gateway-go/pkg/database"
	"ledungcobra/gateway-go/pkg/middlewares"
	"ledungcobra/gateway-go/pkg/routes"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	server *fiber.App
	db     *database.SQLConnector
	config *config.Config
}

func NewServer() interfaces.IServer {
	return &App{}
}

func (a *App) Initialize() error {
	log.Println("Initializing server...")
	a.config = config.New()
	a.setupDatabase(a.config)
	a.setupWebServer()
	return nil
}

func (a *App) setupDatabase(config *config.Config) {
	log.Println("Setting up database")
	a.db = database.NewSQLConnector(config.SqlDsn)
	a.db.Connect()
	a.db.MigrateModels()
	log.Println("Setting up database success")
}

func (a *App) setupWebServer() {
	log.Println("Setting up web server")
	a.server = fiber.New(fiber.Config{
		AppName: "Localhost",
	})
	middlewares.SetBeforeMiddlewares(a.server)
	routes.SetUpRoutes(a.db, a.server, app_service.NewNotificationService())
	middlewares.SetAfterMiddlewares(a.server)
}

func (a *App) Listen() error {
	log.Print("Listening on port " + a.config.ServerPort + "...")
	return a.server.Listen(":" + a.config.ServerPort)
}

func (a *App) Close() error {
	log.Println("Closing server...")
	return nil
}
