package handlers

import (
	"beacon.silali.com/internal/api/core"
	"beacon.silali.com/internal/api/dtos"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetWorkspaces(c echo.Context, app *core.AppContext) error {
	user, err := app.CurrentUser(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("%s", err))
	}

	workspaces, workspaceErr := app.Models.Workspace.GetUserWorkspaces(user.ID)

	if workspaceErr != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("%s", workspaceErr))
	}

	return c.JSON(http.StatusOK, workspaces)
}

func CreateWorkspace(c echo.Context, app *core.AppContext) error {
	request := new(dtos.CreateWorkspaceRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Code:    dtos.ErrorCodeUnsupportedRequest,
			Message: fmt.Sprintf("%s", err),
			Details: err,
		})
	}

	user, err := app.CurrentUser(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("%s", err))
	}

	workspaces, workspaceErr := app.Models.Workspace.CreateWorkspace(user.ID, request.Name)
	if workspaceErr != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("%s", err))
	}

	return c.JSON(http.StatusCreated, workspaces)
}
