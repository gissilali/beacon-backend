package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (app *application) getServersHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, envelope{"data": make([]int, 0)})
}
