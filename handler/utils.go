package handler

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo/v4"
)

func respondJSON(ctx echo.Context, statusCode int, data interface{}) error {

	var response []byte
	var err error
	if ctx.Echo().Debug {
		response, err = json.MarshalIndent(data, "", " ")
	} else {
		response, err = json.Marshal(data)
	}
	if err != nil {
		return HandleError(ctx, err)
	}

	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return respond(ctx, statusCode, response)
}

func respond(ctx echo.Context, statusCode int, data []byte) error {

	res := ctx.Response()

	res.Status = statusCode
	res.WriteHeader(statusCode)

	_, err := ctx.Response().Write((data))
	return HandleError(ctx, err)
}

func HandleError(ctx echo.Context, err error) error {

	if err == nil {
		return nil
	}

	fmt.Println("ERROR: ", err)
	fmt.Println("ERROR.ERROR(): ", err.Error())

	return err
}
