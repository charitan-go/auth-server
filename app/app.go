package app

import (
	"log"

	"github.com/charitan-go/auth-server/external/profile"
	"github.com/charitan-go/auth-server/internal/auth"
	"github.com/charitan-go/auth-server/pkg/database"
	"github.com/charitan-go/auth-server/rabbitmq"
	"github.com/charitan-go/auth-server/rest"
	"github.com/charitan-go/auth-server/rest/api"

	"go.uber.org/fx"
)

// type App struct {
// 	echo *echo.Echo
//
// 	api *api.Api
// }
//
// func newApp(echo *echo.Echo, api *api.Api) *App {
// 	return &App{
// 		echo: echo,
// 		api:  api,
// 	}
// }
//
// func newEcho() *echo.Echo {
// 	return echo.New()
// }
//
// func (app *App) setupRouting() {
// 	// Health Check
// 	app.echo.GET("/health", app.api.HealthCheck)
//
// 	// Auth
// 	app.echo.POST("/donor/register", app.api.AuthHandler.RegisterDonor)
// }
//
// func Run() {
// 	// Register with service registry
// 	discovery.SetupServiceRegistry()
//
// 	// Connect to db
// 	database.SetupDatabase()
//
// 	// TODO: Setup GRPC Service Server
//
// 	fx.New(
// 		profile.ProfileModule,
// 		auth.AuthModule,
// 		fx.Provide(
// 			newApp,
// 			newEcho,
// 			api.NewApi,
// 		),
//
// 		fx.Invoke(func(app *App) {
// 			app.setupRouting()
//
// 			// go app.echo.Start(":8090")
// 			app.echo.Start(":8090")
// 			log.Println("Server started at http://localhost:8090")
// 		}),
// 	).Run()
// }

// Run both servers concurrently
func runServers(restSrv *rest.RestServer, rabbitmqSrv *rabbitmq.RabbitmqServer) {
	log.Println("In invoke")

	// Start REST server
	go func() {
		log.Println("In goroutine of rest")
		restSrv.Run()
	}()

	// Start RabbitMQ server
	go func() {
		log.Println("In goroutine of rabbitmq")
		rabbitmqSrv.Run()
	}()

	// Start gRPC server
	// go func() {
	// 	log.Println("In goroutine of grpc")
	// 	grpcSrv.Run()
	// }()
}

func Run() {
	// Connect to db
	database.SetupDatabase()

	fx.New(
		auth.AuthModule,
		profile.ProfileModule,
		rabbitmq.RabbitmqModule,
		fx.Provide(
			rest.NewRestServer,
			rest.NewEcho,
			api.NewApi,
		),
		fx.Invoke(runServers),
	).Run()
}
