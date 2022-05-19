package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func StartEC2() func(ctx echo.Context) error {
	return func(ctx echo.Context) error {

		return respondJSON(ctx, http.StatusOK, "Aloha!")
	}
}