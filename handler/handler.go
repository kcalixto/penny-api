package handler

import (
	"fmt"
	"net/http"

	"github.com/kcalixto/penny-api/constants"
	"github.com/kcalixto/penny-api/domain/contract"
	"github.com/kcalixto/penny-api/viewmodel"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/labstack/echo/v4"
)

func GetAllInstances() func(ctx echo.Context) error {
	return func(ctx echo.Context) error {

		svc := ec2.New(session.Must(session.NewSession(&aws.Config{
			Region: aws.String(constants.REGION),
		})))
		input := &ec2.DescribeInstancesInput{}

		result, err := svc.DescribeInstances(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				HandleError(ctx, err)
			}
		}

		return respondJSON(ctx, http.StatusOK, viewmodel.ParseDescribeInstancesOutput(result))
	}
}

func GetAllInstancesDetailed() func(ctx echo.Context) error {
	return func(ctx echo.Context) error {

		svc := ec2.New(session.Must(session.NewSession(&aws.Config{
			Region: aws.String(constants.REGION),
		})))
		input := &ec2.DescribeInstancesInput{}

		result, err := svc.DescribeInstances(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				fmt.Println(err.Error())
			}
		}

		return respondJSON(ctx, http.StatusOK, result)
	}
}

func StartEC2(instanceService contract.InstanceService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		context := getContext()

		instanceID := ctx.Param("id")

		result, err := instanceService.StartEC2(context, instanceID)
		if err != nil {
			return HandleError(ctx, err)
		}

		return respondJSON(ctx, http.StatusOK, result)
	}
}

func StopEC2(instanceService contract.InstanceService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		context := getContext()

		instanceID := ctx.Param("id")

		result, err := instanceService.StopEC2(context, instanceID)
		if err != nil {
			return HandleError(ctx, err)
		}

		return respondJSON(ctx, http.StatusOK, result)
	}
}
