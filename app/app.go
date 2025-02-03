package app

import (
	"fmt"

	"github.com/charitan-go/auth-server/api"
	"github.com/charitan-go/auth-server/domain/auth"
	"github.com/charitan-go/auth-server/pkg/database"
	"github.com/charitan-go/auth-server/pkg/discovery"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type App struct {
	echo *echo.Echo

	api *api.Api
}

func newApp(echo *echo.Echo, api *api.Api) *App {
	return &App{
		echo: echo,
		api:  api,
	}
}

func newEcho() *echo.Echo {
	return echo.New()
}

func (app *App) setupRouting() {
	// Health Check
	app.echo.GET("/health", app.api.HealthCheck)

	// Auth
	app.echo.POST("/donor/register", app.api.AuthHandler.RegisterDonor)
}

func (app *App) setup() {
	// Register with service registry
	discovery.SetupServiceRegistry()

	// Connect to db
	database.SetupDatabase()

	// Setup routing
	app.setupRouting()
}

func Run() {
	fx.New(
		auth.AuthModule,
		fx.Provide(
			newApp,
			newEcho,
			api.NewApi,
		),

		fx.Invoke(func(app *App) {
			app.setup()

			go app.echo.Start(":8090")
			fmt.Println("Server started at http://localhost:8090")
		}),
	).Run()
}
