package handlers

import (
	"beacon.silali.com/internal/api/dtos"
	"github.com/labstack/echo/v4"
)

func respondWithError(c echo.Context, status int, message string, err error) error {
	return c.JSON(status, dtos.ErrorResponse{
		Code:    message,
		Message: err.Error(),
		Details: err,
	})
}
