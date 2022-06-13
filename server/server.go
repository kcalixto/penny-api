package server

import (
	"github.com/kcalixto/penny-api/application"
	"github.com/kcalixto/penny-api/handler"

	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
)

func CreateLambda(app application.App) *echoadapter.EchoLambda {
	server := CreateServer(app)
	return echoadapter.New(server)
}

func CreateServer(app application.App) *echo.Echo {
	server := newServer()

	server.GET("/instances", handler.GetAllInstances())
	server.GET("/instances/detail", handler.GetAllInstancesDetailed())
	server.GET("/start/:id", handler.StartEC2(app.Services().Instance()))
	server.GET("/stop/:id", handler.StopEC2(app.Services().Instance()))

	return server
}

func newServer() *echo.Echo {
	e := echo.New()

	//middlewares & stuff goes here

	return e
}
