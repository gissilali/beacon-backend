package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *application) healthcheckHandler(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf(`
		Status: available
		Environment: %s
		Version: %s
	`, app.config.env, app.version))
}
