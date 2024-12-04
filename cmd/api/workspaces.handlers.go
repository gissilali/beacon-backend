package main

import (
	"beacon.silali.com/internal/data"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (app *application) getWorkspacesHandler(c echo.Context) error {
	user, err := app.currentUser(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("%s", err))
	}

	workspaces, workspaceErr := app.models.Workspace.GetUserWorkspaces(user.ID)

	if workspaceErr != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("%s", workspaceErr))
	}

	return c.JSON(http.StatusOK, workspaces)
}

func (app *application) createWorkspaceHandler(c echo.Context) error {
	request := new(data.CreateWorkspaceRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorEnvelope{
			Code:    ErrorCodeUnsupportedRequest,
			Message: fmt.Sprintf("%s", err),
			Details: err,
		})
	}

	user, err := app.currentUser(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("%s", err))
	}

	workspaces, workspaceErr := app.models.Workspace.CreateWorkspace(user.ID, request.Name)
	if workspaceErr != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("%s", err))
	}

	return c.JSON(http.StatusCreated, workspaces)
}
