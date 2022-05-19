package server

import (
	"net/http"
	"penny-api/handler"

	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
)

func startServer() {
	e := echo.New()
	e.GET("/start", func(context echo.Context) error {
		return context.String(http.StatusOK, "Olá!")
	})

	e.Start(":5000")
}

func CreateLambda() *echoadapter.EchoLambda {
	server := CreateServer()
	return echoadapter.New(server)
}

func CreateServer() *echo.Echo {
	server := newServer()

	server.GET("/start", handler.StartEC2())

	return server
}

func newServer() *echo.Echo {
	e := echo.New()

	//middlewares & stuff goes here

	return e
}
