package server

import (
	"github.com/kcalixto/penny-api/handler"

	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
)

func CreateLambda() *echoadapter.EchoLambda {
	server := CreateServer()
	return echoadapter.New(server)
}

func CreateServer() *echo.Echo {
	server := newServer()

	server.GET("/instances", handler.GetAllInstances())
	server.GET("/instances/detail", handler.GetAllInstancesDetailed())
	server.GET("/start/:id", handler.StartEC2())

	return server
}

func newServer() *echo.Echo {
	e := echo.New()

	//middlewares & stuff goes here

	return e
}
