package main

import (
	"context"
	"fmt"
	"os"
	"penny-api/server"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
)

var PORT = "5000"

var echoLambda *echoadapter.EchoLambda

func init() {
	echoLambda = server.CreateLambda()
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambdaDefaultEnvPort := os.Getenv("_LAMBDA_SERVER_PORT")
	lambdaDefaultEnvRuntime := os.Getenv("AWS_LAMBDA_RUNTIME_API")

	if lambdaDefaultEnvPort == "" && lambdaDefaultEnvRuntime == "" {
		server := server.CreateServer()
		err := server.Start(":" + PORT)
		if err != nil {
			fmt.Print("env_error: ", err)
		}
	} else {
		lambda.Start(Handler)
	}
}
