package app

import (
	"log"

	"ledungcobra/gateway-go/pkg/interfaces"

	"ledungcobra/gateway-go/pkg/config"
	"ledungcobra/gateway-go/pkg/database"
	"ledungcobra/gateway-go/pkg/middlewares"
	"ledungcobra/gateway-go/pkg/routes"

	"github.com/gofiber/fiber/v2"
)

func NewServer() interfaces.IServer {
	return &App{}
}

type App struct {
	server *fiber.App
	db     *database.SQLDBManager
	config *config.Config
}

// Initialize
func (a *App) Initialize() error {
	log.Println("Initializing server...")
	config.Init()
	a.config = config.Cfg
	a.setupDatabase(a.config)
	a.setupWebServer()
	return nil
}

func (a *App) setupDatabase(config *config.Config) {
	log.Println("Setting up database")
	a.db = database.NewSQLDatabase(config.SqlDsn)
	a.db.Connect()
	a.db.MigrateModels()
	log.Println("Setting up database success")
}

func (a *App) setupWebServer() {
	log.Println("Setting up web server")
	a.server = fiber.New(fiber.Config{
		AppName: "Localhost",
	})
	routes.SetUpRoutes(a.db, a.server, a.config)
	middlewares.SetUpMiddlewares(a.server)
}

func (a *App) Listen() error {
	log.Print("Listening on port " + a.config.ServerPort + "...")
	return a.server.Listen(":" + a.config.ServerPort)
}

func (*App) Stop() error {
	return nil
}
