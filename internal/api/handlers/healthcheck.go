package handlers

import (
	"beacon.silali.com/internal/api/core"
	"github.com/labstack/echo/v4"
	"net/http"
)

type envelope map[string]interface{}

func HealthCheck(c echo.Context, app *core.AppContext) error {
	return c.JSON(http.StatusOK, envelope{
		"status":  "Ok",
		"version": app.Version,
	})
}
